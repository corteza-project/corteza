package rbac

import "github.com/cortezaproject/corteza-server/pkg/slice"

// Merge applies new rules (changes) to existing set and mark all changes as dirty
func (set RuleSet) Merge(rules ...*Rule) (out RuleSet) {
	var (
		o    int
		olen = len(set)
	)

	if olen == 0 {
		// Nothing exists yet, mark all as dirty
		for r := range rules {
			rules[r].dirty = true
		}

		return rules
	} else {
		out = set

	newRules:
		for _, rule := range rules {
			// Never go beyond the last old rule (olen)
			for o = 0; o < olen; o++ {
				if out[o].Equals(rule) {
					out[o].dirty = out[o].Access != rule.Access
					out[o].Access = rule.Access

					// only one rule can match so proceed with next new rule
					continue newRules
				}
			}

			// none of the old rules matched, append
			var c = *rule
			c.dirty = true

			out = append(out, &c)
		}

	}

	return
}

// Dirty returns list of changed (Dirty==true) and deleted (Access==Inherit) rules
func (set RuleSet) Dirty() (inherited, rest RuleSet) {
	inherited, rest = RuleSet{}, RuleSet{}

	for _, r := range set {
		var c = *r
		if r.Access == Inherit {
			inherited = append(inherited, &c)
		} else if r.dirty {
			rest = append(rest, &c)
		}
	}

	return
}

// reset dirty flag
func (set RuleSet) Clear() {
	_ = set.Walk(func(rule *Rule) error {
		rule.dirty = false
		return nil
	})
}

// Missing compares cmp with existing set
// and returns rules that exists in set but not in cmp
func (set RuleSet) Diff(cmp RuleSet) RuleSet {
	diff := RuleSet{}
base:
	for _, s := range set {
		for _, c := range cmp {
			if c.Equals(s) {
				continue base
			}
		}

		diff = append(diff, s)
	}

	return diff
}

// Roles returns list of unique id of all roles in the rule set
func (set RuleSet) Roles() []uint64 {
	roles := make([]uint64, 0)
	for _, r := range set {
		if slice.HasUint64(roles, r.RoleID) {
			continue
		}

		roles = append(roles, r.RoleID)
	}

	return roles
}

func (set RuleSet) ByResource(res Resource) RuleSet {
	out, _ := set.Filter(func(r *Rule) (bool, error) {
		return res == r.Resource, nil
	})
	return out
}

func (set RuleSet) AllAllows() RuleSet {
	return set.ByAccess(Allow)
}

func (set RuleSet) AllDenies() RuleSet {
	return set.ByAccess(Deny)
}

func (set RuleSet) ByAccess(a Access) RuleSet {
	out, _ := set.Filter(func(r *Rule) (bool, error) {
		return a == r.Access, nil
	})
	return out
}

func (set RuleSet) ByRole(roleID uint64) RuleSet {
	out, _ := set.Filter(func(r *Rule) (bool, error) {
		return roleID == r.RoleID, nil
	})
	return out
}

// Significant roles (sigRoles) returns two list of significant roles.
//
// First slice are roles that are allowed to perform an operation on a specific resource (directly or indirectly),
// 2nd slice are roles that are denied the op.
func (set RuleSet) sigRoles(res Resource, op Operation) (aRR, dRR []uint64) {
	if !res.IsAppendable() {
		// nothing to do here, we need a direct resource id (level=0)
		return
	}

	const (
		dirRes = 0
		indRes = 1
	)

	var (
		rr = map[Access]map[int][]uint64{
			Allow: {
				dirRes: make([]uint64, 0),
				indRes: make([]uint64, 0),
			},
			Deny: {
				dirRes: make([]uint64, 0),
				indRes: make([]uint64, 0),
			},
		}
	)

	// Extract all relevant rules (by op and resource) and group them by
	// access and distance (rules for direct resources and rules for indirect resources)
	for _, r := range set {
		if r.Operation != op {
			continue
		}

		if r.Resource == res {
			// direct rules
			rr[r.Access][dirRes] = append(rr[r.Access][dirRes], r.RoleID)
		}

		if r.Resource.IsAppendable() && r.Resource.TrimID() == res.TrimID() {
			// rules on all resources of this type
			rr[r.Access][indRes] = append(rr[r.Access][indRes], r.RoleID)
		}
	}

	// Process all extracted roles and make sure that are filtered and ordered by relevance:
	//  1. list of roles with denied operation directly on the resource
	dRoles := slice.ToUint64BoolMap(rr[Deny][dirRes])

	//  2. list of roles with allowed operation directly on the resource
	aRoles := make(map[uint64]bool)
	for _, r := range rr[Allow][dirRes] {
		aRoles[r] = !dRoles[r]
	}

	//  3. list of roles with denied operation indirectly on the resource
	for _, r := range rr[Deny][indRes] {
		dRoles[r] = true
	}

	//  4. list of roles with allowed operation indirectly on the resource
	for _, r := range rr[Allow][indRes] {
		aRoles[r] = !dRoles[r]
	}

	for r, chk := range aRoles {
		if chk {
			aRR = append(aRR, r)
		}
	}

	for r, chk := range dRoles {
		if chk {
			dRR = append(dRR, r)
		}
	}

	return
}
