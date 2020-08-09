// Code generated by entc, DO NOT EDIT.

package ent

import (
	"flashcards-backend/ent/cardlog"
	"flashcards-backend/ent/cardschedule"
	"flashcards-backend/ent/schema"
	"flashcards-backend/ent/user"
	"flashcards-backend/ent/word"
	"time"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	cardlogMixin := schema.CardLog{}.Mixin()
	cardlogMixinFields0 := cardlogMixin[0].Fields()
	cardlogFields := schema.CardLog{}.Fields()
	_ = cardlogFields
	// cardlogDescCreateTime is the schema descriptor for create_time field.
	cardlogDescCreateTime := cardlogMixinFields0[0].Descriptor()
	// cardlog.DefaultCreateTime holds the default value on creation for the create_time field.
	cardlog.DefaultCreateTime = cardlogDescCreateTime.Default.(func() time.Time)
	cardscheduleMixin := schema.CardSchedule{}.Mixin()
	cardscheduleMixinFields0 := cardscheduleMixin[0].Fields()
	cardscheduleFields := schema.CardSchedule{}.Fields()
	_ = cardscheduleFields
	// cardscheduleDescCreateTime is the schema descriptor for create_time field.
	cardscheduleDescCreateTime := cardscheduleMixinFields0[0].Descriptor()
	// cardschedule.DefaultCreateTime holds the default value on creation for the create_time field.
	cardschedule.DefaultCreateTime = cardscheduleDescCreateTime.Default.(func() time.Time)
	// cardscheduleDescReviewed is the schema descriptor for reviewed field.
	cardscheduleDescReviewed := cardscheduleFields[1].Descriptor()
	// cardschedule.DefaultReviewed holds the default value on creation for the reviewed field.
	cardschedule.DefaultReviewed = cardscheduleDescReviewed.Default.(bool)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userMixinFields0[0].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userMixinFields0[1].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[0].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = func() func(string) error {
		validators := userDescEmail.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(email string) error {
			for _, fn := range fns {
				if err := fn(email); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	wordMixin := schema.Word{}.Mixin()
	wordMixinFields0 := wordMixin[0].Fields()
	wordFields := schema.Word{}.Fields()
	_ = wordFields
	// wordDescCreateTime is the schema descriptor for create_time field.
	wordDescCreateTime := wordMixinFields0[0].Descriptor()
	// word.DefaultCreateTime holds the default value on creation for the create_time field.
	word.DefaultCreateTime = wordDescCreateTime.Default.(func() time.Time)
	// wordDescUpdateTime is the schema descriptor for update_time field.
	wordDescUpdateTime := wordMixinFields0[1].Descriptor()
	// word.DefaultUpdateTime holds the default value on creation for the update_time field.
	word.DefaultUpdateTime = wordDescUpdateTime.Default.(func() time.Time)
	// word.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	word.UpdateDefaultUpdateTime = wordDescUpdateTime.UpdateDefault.(func() time.Time)
	// wordDescLang1 is the schema descriptor for lang1 field.
	wordDescLang1 := wordFields[0].Descriptor()
	// word.Lang1Validator is a validator for the "lang1" field. It is called by the builders before save.
	word.Lang1Validator = func() func(string) error {
		validators := wordDescLang1.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(lang1 string) error {
			for _, fn := range fns {
				if err := fn(lang1); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// wordDescLang2 is the schema descriptor for lang2 field.
	wordDescLang2 := wordFields[1].Descriptor()
	// word.Lang2Validator is a validator for the "lang2" field. It is called by the builders before save.
	word.Lang2Validator = func() func(string) error {
		validators := wordDescLang2.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(lang2 string) error {
			for _, fn := range fns {
				if err := fn(lang2); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// wordDescWord1 is the schema descriptor for word1 field.
	wordDescWord1 := wordFields[2].Descriptor()
	// word.Word1Validator is a validator for the "word1" field. It is called by the builders before save.
	word.Word1Validator = func() func(string) error {
		validators := wordDescWord1.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(word1 string) error {
			for _, fn := range fns {
				if err := fn(word1); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// wordDescWord2 is the schema descriptor for word2 field.
	wordDescWord2 := wordFields[3].Descriptor()
	// word.Word2Validator is a validator for the "word2" field. It is called by the builders before save.
	word.Word2Validator = func() func(string) error {
		validators := wordDescWord2.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(word2 string) error {
			for _, fn := range fns {
				if err := fn(word2); err != nil {
					return err
				}
			}
			return nil
		}
	}()
}
