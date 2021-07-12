package types

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// system/types/types.yaml

type (

	// ApigwFunctionSet slice of ApigwFunction
	//
	// This type is auto-generated.
	ApigwFunctionSet []*ApigwFunction

	// ApigwRouteSet slice of ApigwRoute
	//
	// This type is auto-generated.
	ApigwRouteSet []*ApigwRoute

	// ApplicationSet slice of Application
	//
	// This type is auto-generated.
	ApplicationSet []*Application

	// AttachmentSet slice of Attachment
	//
	// This type is auto-generated.
	AttachmentSet []*Attachment

	// AuthClientSet slice of AuthClient
	//
	// This type is auto-generated.
	AuthClientSet []*AuthClient

	// AuthConfirmedClientSet slice of AuthConfirmedClient
	//
	// This type is auto-generated.
	AuthConfirmedClientSet []*AuthConfirmedClient

	// AuthOa2tokenSet slice of AuthOa2token
	//
	// This type is auto-generated.
	AuthOa2tokenSet []*AuthOa2token

	// AuthSessionSet slice of AuthSession
	//
	// This type is auto-generated.
	AuthSessionSet []*AuthSession

	// CredentialsSet slice of Credentials
	//
	// This type is auto-generated.
	CredentialsSet []*Credentials

	// ReminderSet slice of Reminder
	//
	// This type is auto-generated.
	ReminderSet []*Reminder

	// RoleSet slice of Role
	//
	// This type is auto-generated.
	RoleSet []*Role

	// RoleMemberSet slice of RoleMember
	//
	// This type is auto-generated.
	RoleMemberSet []*RoleMember

	// SettingValueSet slice of SettingValue
	//
	// This type is auto-generated.
	SettingValueSet []*SettingValue

	// TemplateSet slice of Template
	//
	// This type is auto-generated.
	TemplateSet []*Template

	// UserSet slice of User
	//
	// This type is auto-generated.
	UserSet []*User
)

// Walk iterates through every slice item and calls w(ApigwFunction) err
//
// This function is auto-generated.
func (set ApigwFunctionSet) Walk(w func(*ApigwFunction) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(ApigwFunction) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set ApigwFunctionSet) Filter(f func(*ApigwFunction) (bool, error)) (out ApigwFunctionSet, err error) {
	var ok bool
	out = ApigwFunctionSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// FindByID finds items from slice by its ID property
//
// This function is auto-generated.
func (set ApigwFunctionSet) FindByID(ID uint64) *ApigwFunction {
	for i := range set {
		if set[i].ID == ID {
			return set[i]
		}
	}

	return nil
}

// IDs returns a slice of uint64s from all items in the set
//
// This function is auto-generated.
func (set ApigwFunctionSet) IDs() (IDs []uint64) {
	IDs = make([]uint64, len(set))

	for i := range set {
		IDs[i] = set[i].ID
	}

	return
}

// Walk iterates through every slice item and calls w(ApigwRoute) err
//
// This function is auto-generated.
func (set ApigwRouteSet) Walk(w func(*ApigwRoute) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(ApigwRoute) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set ApigwRouteSet) Filter(f func(*ApigwRoute) (bool, error)) (out ApigwRouteSet, err error) {
	var ok bool
	out = ApigwRouteSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// FindByID finds items from slice by its ID property
//
// This function is auto-generated.
func (set ApigwRouteSet) FindByID(ID uint64) *ApigwRoute {
	for i := range set {
		if set[i].ID == ID {
			return set[i]
		}
	}

	return nil
}

// IDs returns a slice of uint64s from all items in the set
//
// This function is auto-generated.
func (set ApigwRouteSet) IDs() (IDs []uint64) {
	IDs = make([]uint64, len(set))

	for i := range set {
		IDs[i] = set[i].ID
	}

	return
}

// Walk iterates through every slice item and calls w(Application) err
//
// This function is auto-generated.
func (set ApplicationSet) Walk(w func(*Application) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(Application) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set ApplicationSet) Filter(f func(*Application) (bool, error)) (out ApplicationSet, err error) {
	var ok bool
	out = ApplicationSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// FindByID finds items from slice by its ID property
//
// This function is auto-generated.
func (set ApplicationSet) FindByID(ID uint64) *Application {
	for i := range set {
		if set[i].ID == ID {
			return set[i]
		}
	}

	return nil
}

// IDs returns a slice of uint64s from all items in the set
//
// This function is auto-generated.
func (set ApplicationSet) IDs() (IDs []uint64) {
	IDs = make([]uint64, len(set))

	for i := range set {
		IDs[i] = set[i].ID
	}

	return
}

// Walk iterates through every slice item and calls w(Attachment) err
//
// This function is auto-generated.
func (set AttachmentSet) Walk(w func(*Attachment) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(Attachment) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set AttachmentSet) Filter(f func(*Attachment) (bool, error)) (out AttachmentSet, err error) {
	var ok bool
	out = AttachmentSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// FindByID finds items from slice by its ID property
//
// This function is auto-generated.
func (set AttachmentSet) FindByID(ID uint64) *Attachment {
	for i := range set {
		if set[i].ID == ID {
			return set[i]
		}
	}

	return nil
}

// IDs returns a slice of uint64s from all items in the set
//
// This function is auto-generated.
func (set AttachmentSet) IDs() (IDs []uint64) {
	IDs = make([]uint64, len(set))

	for i := range set {
		IDs[i] = set[i].ID
	}

	return
}

// Walk iterates through every slice item and calls w(AuthClient) err
//
// This function is auto-generated.
func (set AuthClientSet) Walk(w func(*AuthClient) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(AuthClient) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set AuthClientSet) Filter(f func(*AuthClient) (bool, error)) (out AuthClientSet, err error) {
	var ok bool
	out = AuthClientSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// FindByID finds items from slice by its ID property
//
// This function is auto-generated.
func (set AuthClientSet) FindByID(ID uint64) *AuthClient {
	for i := range set {
		if set[i].ID == ID {
			return set[i]
		}
	}

	return nil
}

// IDs returns a slice of uint64s from all items in the set
//
// This function is auto-generated.
func (set AuthClientSet) IDs() (IDs []uint64) {
	IDs = make([]uint64, len(set))

	for i := range set {
		IDs[i] = set[i].ID
	}

	return
}

// Walk iterates through every slice item and calls w(AuthConfirmedClient) err
//
// This function is auto-generated.
func (set AuthConfirmedClientSet) Walk(w func(*AuthConfirmedClient) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(AuthConfirmedClient) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set AuthConfirmedClientSet) Filter(f func(*AuthConfirmedClient) (bool, error)) (out AuthConfirmedClientSet, err error) {
	var ok bool
	out = AuthConfirmedClientSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// Walk iterates through every slice item and calls w(AuthOa2token) err
//
// This function is auto-generated.
func (set AuthOa2tokenSet) Walk(w func(*AuthOa2token) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(AuthOa2token) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set AuthOa2tokenSet) Filter(f func(*AuthOa2token) (bool, error)) (out AuthOa2tokenSet, err error) {
	var ok bool
	out = AuthOa2tokenSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// FindByID finds items from slice by its ID property
//
// This function is auto-generated.
func (set AuthOa2tokenSet) FindByID(ID uint64) *AuthOa2token {
	for i := range set {
		if set[i].ID == ID {
			return set[i]
		}
	}

	return nil
}

// IDs returns a slice of uint64s from all items in the set
//
// This function is auto-generated.
func (set AuthOa2tokenSet) IDs() (IDs []uint64) {
	IDs = make([]uint64, len(set))

	for i := range set {
		IDs[i] = set[i].ID
	}

	return
}

// Walk iterates through every slice item and calls w(AuthSession) err
//
// This function is auto-generated.
func (set AuthSessionSet) Walk(w func(*AuthSession) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(AuthSession) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set AuthSessionSet) Filter(f func(*AuthSession) (bool, error)) (out AuthSessionSet, err error) {
	var ok bool
	out = AuthSessionSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// Walk iterates through every slice item and calls w(Credentials) err
//
// This function is auto-generated.
func (set CredentialsSet) Walk(w func(*Credentials) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(Credentials) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set CredentialsSet) Filter(f func(*Credentials) (bool, error)) (out CredentialsSet, err error) {
	var ok bool
	out = CredentialsSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// FindByID finds items from slice by its ID property
//
// This function is auto-generated.
func (set CredentialsSet) FindByID(ID uint64) *Credentials {
	for i := range set {
		if set[i].ID == ID {
			return set[i]
		}
	}

	return nil
}

// IDs returns a slice of uint64s from all items in the set
//
// This function is auto-generated.
func (set CredentialsSet) IDs() (IDs []uint64) {
	IDs = make([]uint64, len(set))

	for i := range set {
		IDs[i] = set[i].ID
	}

	return
}

// Walk iterates through every slice item and calls w(Reminder) err
//
// This function is auto-generated.
func (set ReminderSet) Walk(w func(*Reminder) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(Reminder) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set ReminderSet) Filter(f func(*Reminder) (bool, error)) (out ReminderSet, err error) {
	var ok bool
	out = ReminderSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// FindByID finds items from slice by its ID property
//
// This function is auto-generated.
func (set ReminderSet) FindByID(ID uint64) *Reminder {
	for i := range set {
		if set[i].ID == ID {
			return set[i]
		}
	}

	return nil
}

// IDs returns a slice of uint64s from all items in the set
//
// This function is auto-generated.
func (set ReminderSet) IDs() (IDs []uint64) {
	IDs = make([]uint64, len(set))

	for i := range set {
		IDs[i] = set[i].ID
	}

	return
}

// Walk iterates through every slice item and calls w(Role) err
//
// This function is auto-generated.
func (set RoleSet) Walk(w func(*Role) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(Role) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set RoleSet) Filter(f func(*Role) (bool, error)) (out RoleSet, err error) {
	var ok bool
	out = RoleSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// FindByID finds items from slice by its ID property
//
// This function is auto-generated.
func (set RoleSet) FindByID(ID uint64) *Role {
	for i := range set {
		if set[i].ID == ID {
			return set[i]
		}
	}

	return nil
}

// IDs returns a slice of uint64s from all items in the set
//
// This function is auto-generated.
func (set RoleSet) IDs() (IDs []uint64) {
	IDs = make([]uint64, len(set))

	for i := range set {
		IDs[i] = set[i].ID
	}

	return
}

// Walk iterates through every slice item and calls w(RoleMember) err
//
// This function is auto-generated.
func (set RoleMemberSet) Walk(w func(*RoleMember) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(RoleMember) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set RoleMemberSet) Filter(f func(*RoleMember) (bool, error)) (out RoleMemberSet, err error) {
	var ok bool
	out = RoleMemberSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// Walk iterates through every slice item and calls w(SettingValue) err
//
// This function is auto-generated.
func (set SettingValueSet) Walk(w func(*SettingValue) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(SettingValue) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set SettingValueSet) Filter(f func(*SettingValue) (bool, error)) (out SettingValueSet, err error) {
	var ok bool
	out = SettingValueSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// Walk iterates through every slice item and calls w(Template) err
//
// This function is auto-generated.
func (set TemplateSet) Walk(w func(*Template) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(Template) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set TemplateSet) Filter(f func(*Template) (bool, error)) (out TemplateSet, err error) {
	var ok bool
	out = TemplateSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// FindByID finds items from slice by its ID property
//
// This function is auto-generated.
func (set TemplateSet) FindByID(ID uint64) *Template {
	for i := range set {
		if set[i].ID == ID {
			return set[i]
		}
	}

	return nil
}

// IDs returns a slice of uint64s from all items in the set
//
// This function is auto-generated.
func (set TemplateSet) IDs() (IDs []uint64) {
	IDs = make([]uint64, len(set))

	for i := range set {
		IDs[i] = set[i].ID
	}

	return
}

// Walk iterates through every slice item and calls w(User) err
//
// This function is auto-generated.
func (set UserSet) Walk(w func(*User) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(User) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set UserSet) Filter(f func(*User) (bool, error)) (out UserSet, err error) {
	var ok bool
	out = UserSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// FindByID finds items from slice by its ID property
//
// This function is auto-generated.
func (set UserSet) FindByID(ID uint64) *User {
	for i := range set {
		if set[i].ID == ID {
			return set[i]
		}
	}

	return nil
}

// IDs returns a slice of uint64s from all items in the set
//
// This function is auto-generated.
func (set UserSet) IDs() (IDs []uint64) {
	IDs = make([]uint64, len(set))

	for i := range set {
		IDs[i] = set[i].ID
	}

	return
}
