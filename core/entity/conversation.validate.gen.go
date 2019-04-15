// this file was generated by protoc-gen-gotemplate

package entity

import (
	"github.com/pkg/errors"

	"berty.tech/core/pkg/errorcodes"
	"berty.tech/core/pkg/validate/validator"
)

var (
	_ = errors.New
	_ = validator.IsContactKey
	_ = errorcodes.IsSubCode
)

func (m *Conversation) Validate() error {
	if m == nil {
		return nil
	}

	// handling field: ID - name:"id" number:1 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"id" options:<53004:1 []:1 []:"ID" 65006:"gorm:\"primary_key\"" >  (is_contact_key=false, defined_only=false, min_len=1, max_len=0, skip=false, required=false, min_items=0, max_items=0)
	if len(m.GetID()) < 1 {
		return errors.New("ID must be longer than 1")
	}

	// handling field: CreatedAt - name:"created_at" number:2 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".google.protobuf.Timestamp" json_name:"createdAt" options:<65001:0 65010:1 >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: CreatedAt")
		}
	}

	// handling field: UpdatedAt - name:"updated_at" number:3 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".google.protobuf.Timestamp" json_name:"updatedAt" options:<65001:0 65010:1 >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: UpdatedAt")
		}
	}

	// handling field: ReadAt - name:"read_at" number:4 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".google.protobuf.Timestamp" json_name:"readAt" options:<65001:0 65010:1 >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetReadAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: ReadAt")
		}
	}

	// handling field: WroteAt - name:"wrote_at" number:5 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".google.protobuf.Timestamp" json_name:"wroteAt" options:<65001:0 65010:1 >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetWroteAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: WroteAt")
		}
	}

	// handling field: Title - name:"title" number:20 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"title"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: Topic - name:"topic" number:21 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"topic"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: Infos - name:"infos" number:22 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"infos"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: Kind - name:"kind" number:23 label:LABEL_OPTIONAL type:TYPE_ENUM type_name:".berty.entity.Conversation.Kind" json_name:"kind"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: Members - name:"members" number:100 label:LABEL_REPEATED type:TYPE_MESSAGE type_name:".berty.entity.ConversationMember" json_name:"members" options:<65006:"gorm:\"foreignkey:ConversationID;association_foreignkey:ID;save_associations:true\"" >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetMembers()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: Members")
		}
	}

	return nil
}
func (m *ConversationMember) Validate() error {
	if m == nil {
		return nil
	}

	// handling field: ID - name:"id" number:1 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"id" options:<53004:1 []:1 []:"ID" 65006:"gorm:\"primary_key\"" >  (is_contact_key=false, defined_only=false, min_len=1, max_len=0, skip=false, required=false, min_items=0, max_items=0)
	if len(m.GetID()) < 1 {
		return errors.New("ID must be longer than 1")
	}

	// handling field: CreatedAt - name:"created_at" number:2 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".google.protobuf.Timestamp" json_name:"createdAt" options:<65001:0 65010:1 >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: CreatedAt")
		}
	}

	// handling field: UpdatedAt - name:"updated_at" number:3 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".google.protobuf.Timestamp" json_name:"updatedAt" options:<65001:0 65010:1 >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: UpdatedAt")
		}
	}

	// handling field: ReadAt - name:"read_at" number:4 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".google.protobuf.Timestamp" json_name:"readAt" options:<65001:0 65010:1 >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetReadAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: ReadAt")
		}
	}

	// handling field: WroteAt - name:"wrote_at" number:5 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".google.protobuf.Timestamp" json_name:"wroteAt" options:<65001:0 65010:1 >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetWroteAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: WroteAt")
		}
	}

	// handling field: Status - name:"status" number:10 label:LABEL_OPTIONAL type:TYPE_ENUM type_name:".berty.entity.ConversationMember.Status" json_name:"status"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: Contact - name:"contact" number:100 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".berty.entity.Contact" json_name:"contact" options:<[]:true []:false 65006:"gorm:\"association_autoupdate:false;association_create:true\"" >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=true, min_items=0, max_items=0)
	if m.GetContact() == nil {
		return errors.New("Contact is required")
	}

	if v, ok := interface{}(m.GetContact()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: Contact")
		}
	}

	// handling field: ConversationID - name:"conversation_id" number:101 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"conversationId" options:<53004:1 []:1 []:"ConversationID" >  (is_contact_key=false, defined_only=false, min_len=1, max_len=0, skip=false, required=false, min_items=0, max_items=0)
	if len(m.GetConversationID()) < 1 {
		return errors.New("ConversationID must be longer than 1")
	}

	// handling field: ContactID - name:"contact_id" number:102 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"contactId" options:<53004:1 []:1 []:"ContactID" >  (is_contact_key=false, defined_only=false, min_len=1, max_len=0, skip=false, required=false, min_items=0, max_items=0)
	if len(m.GetContactID()) < 1 {
		return errors.New("ContactID must be longer than 1")
	}
	return nil
}
