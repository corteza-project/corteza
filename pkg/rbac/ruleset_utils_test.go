package rbac

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Test role inheritance
func TestRuleSet_merge(t *testing.T) {
	var (
		req = require.New(t)

		sCases = []struct {
			old RuleSet
			new RuleSet
			del RuleSet
			upd RuleSet
		}{
			{
				RuleSet{AllowRule(role1, resService1, opAccess)},
				RuleSet{AllowRule(role1, resService1, opAccess)},
				RuleSet{},
				RuleSet{},
			},
			{
				RuleSet{AllowRule(role1, resService1, opAccess)},
				RuleSet{DenyRule(role1, resService1, opAccess)},
				RuleSet{},
				RuleSet{DenyRule(role1, resService1, opAccess)},
			},
			{
				RuleSet{AllowRule(role1, resService1, opAccess)},
				RuleSet{InheritRule(role1, resService1, opAccess)},
				RuleSet{InheritRule(role1, resService1, opAccess)},
				RuleSet{},
			},
			{
				RuleSet{AllowRule(role1, resService1, opAccess)},
				RuleSet{AllowRule(role1, resService1, opAccess)},
				RuleSet{},
				RuleSet{},
			},
			{
				RuleSet{
					AllowRule(role1, resService1, opAccess),
					DenyRule(role2, resService1, opAccess),
					DenyRule(EveryoneRoleID, resService2, opAccess),
					AllowRule(role1, resService2, opAccess),
					AllowRule(role2, resThing42, opAccess),
				},
				RuleSet{
					DenyRule(EveryoneRoleID, resThingWc, opAccess),
					AllowRule(role1, resService2, opAccess),
					AllowRule(role1, resThing42, opAccess),
					InheritRule(role2, resThing42, opAccess),
				},
				RuleSet{
					InheritRule(role2, resThing42, opAccess),
				},
				RuleSet{
					// AllowRule(role1, resService1, opAccess),
					// DenyRule(role2, resService1, opAccess),
					// DenyRule(EveryoneRoleID, resService2, opAccess),
					// AllowRule(role1, resService2, opAccess),
					DenyRule(EveryoneRoleID, resThingWc, opAccess),
					AllowRule(role1, resThing42, opAccess),
				},
			},
		}
	)

	for _, sc := range sCases {
		// Apply changed and get update candidates
		mrg := sc.old.Merge(sc.new...)
		del, upd := mrg.Dirty()

		// Clear dirty flag so that we do not confuse DeepEqual
		del.Clear()
		upd.Clear()

		req.Equal(len(sc.del), len(del))
		req.Equal(len(sc.upd), len(upd))
		req.Equal(sc.del, del)
		req.Equal(sc.upd, upd)
	}
}

// Test role inheritance
func TestRuleSet_sigRoles(t *testing.T) {
	const (
		roleA uint64 = 1
		roleB uint64 = 2
		roleC uint64 = 3
		roleD uint64 = 4
		roleE uint64 = 5

		opRead  = Operation("read")
		opWrite = Operation("write")
		resD    = Resource("res:foo:1")
		resI    = Resource("res:foo:*")
	)

	var (
		rr     = func(rr ...uint64) []uint64 { return rr }
		sCases = []struct {
			set RuleSet
			arr []uint64
			drr []uint64
		}{
			{
				RuleSet{
					AllowRule(roleA, resD, opRead),
					DenyRule(roleA, resD, opWrite),
					AllowRule(roleB, resI, opRead),
					DenyRule(roleA, resI, opRead),
					DenyRule(roleC, resI, opRead),
					DenyRule(roleD, resI, opRead),
					DenyRule(roleE, resI, opRead),
					DenyRule(roleE, resI, opWrite),
				},
				rr(roleB),
				rr(roleA, roleC, roleD, roleE),
			},
			{
				RuleSet{
					AllowRule(roleA, resD, opRead),
					DenyRule(roleA, resD, opRead),
				},
				rr(),
				rr(roleA),
			},
			{
				RuleSet{
					AllowRule(roleA, resD, opRead),
					DenyRule(roleA, resD, opRead),
					AllowRule(roleA, resI, opRead),
					DenyRule(roleA, resI, opRead),
				},
				rr(),
				rr(roleA),
			},
		}
	)

	for _, sc := range sCases {
		t.Run("a", func(t *testing.T) {
			req := require.New(t)
			arr, drr := sc.set.sigRoles(resD, opRead)

			req.Equal(sc.arr, arr)
			req.Equal(sc.drr, drr)
		})
	}
}
