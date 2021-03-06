// Code generated by entc, DO NOT EDIT.

package product

import (
	"online-supermarket/controllers/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Sku applies equality check predicate on the "sku" field. It's identical to SkuEQ.
func Sku(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSku), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// Weight applies equality check predicate on the "weight" field. It's identical to WeightEQ.
func Weight(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWeight), v))
	})
}

// Descriptions applies equality check predicate on the "descriptions" field. It's identical to DescriptionsEQ.
func Descriptions(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescriptions), v))
	})
}

// Thumbnail applies equality check predicate on the "thumbnail" field. It's identical to ThumbnailEQ.
func Thumbnail(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThumbnail), v))
	})
}

// CreateDate applies equality check predicate on the "create_date" field. It's identical to CreateDateEQ.
func CreateDate(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateDate), v))
	})
}

// Stock applies equality check predicate on the "stock" field. It's identical to StockEQ.
func Stock(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStock), v))
	})
}

// SkuEQ applies the EQ predicate on the "sku" field.
func SkuEQ(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSku), v))
	})
}

// SkuNEQ applies the NEQ predicate on the "sku" field.
func SkuNEQ(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSku), v))
	})
}

// SkuIn applies the In predicate on the "sku" field.
func SkuIn(vs ...int) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSku), v...))
	})
}

// SkuNotIn applies the NotIn predicate on the "sku" field.
func SkuNotIn(vs ...int) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSku), v...))
	})
}

// SkuGT applies the GT predicate on the "sku" field.
func SkuGT(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSku), v))
	})
}

// SkuGTE applies the GTE predicate on the "sku" field.
func SkuGTE(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSku), v))
	})
}

// SkuLT applies the LT predicate on the "sku" field.
func SkuLT(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSku), v))
	})
}

// SkuLTE applies the LTE predicate on the "sku" field.
func SkuLTE(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSku), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrice), v))
	})
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...float64) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPrice), v...))
	})
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...float64) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPrice), v...))
	})
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrice), v))
	})
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrice), v))
	})
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrice), v))
	})
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrice), v))
	})
}

// WeightEQ applies the EQ predicate on the "weight" field.
func WeightEQ(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWeight), v))
	})
}

// WeightNEQ applies the NEQ predicate on the "weight" field.
func WeightNEQ(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWeight), v))
	})
}

// WeightIn applies the In predicate on the "weight" field.
func WeightIn(vs ...float64) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWeight), v...))
	})
}

// WeightNotIn applies the NotIn predicate on the "weight" field.
func WeightNotIn(vs ...float64) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWeight), v...))
	})
}

// WeightGT applies the GT predicate on the "weight" field.
func WeightGT(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWeight), v))
	})
}

// WeightGTE applies the GTE predicate on the "weight" field.
func WeightGTE(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWeight), v))
	})
}

// WeightLT applies the LT predicate on the "weight" field.
func WeightLT(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWeight), v))
	})
}

// WeightLTE applies the LTE predicate on the "weight" field.
func WeightLTE(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWeight), v))
	})
}

// DescriptionsEQ applies the EQ predicate on the "descriptions" field.
func DescriptionsEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescriptions), v))
	})
}

// DescriptionsNEQ applies the NEQ predicate on the "descriptions" field.
func DescriptionsNEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescriptions), v))
	})
}

// DescriptionsIn applies the In predicate on the "descriptions" field.
func DescriptionsIn(vs ...string) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescriptions), v...))
	})
}

// DescriptionsNotIn applies the NotIn predicate on the "descriptions" field.
func DescriptionsNotIn(vs ...string) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescriptions), v...))
	})
}

// DescriptionsGT applies the GT predicate on the "descriptions" field.
func DescriptionsGT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescriptions), v))
	})
}

// DescriptionsGTE applies the GTE predicate on the "descriptions" field.
func DescriptionsGTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescriptions), v))
	})
}

// DescriptionsLT applies the LT predicate on the "descriptions" field.
func DescriptionsLT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescriptions), v))
	})
}

// DescriptionsLTE applies the LTE predicate on the "descriptions" field.
func DescriptionsLTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescriptions), v))
	})
}

// DescriptionsContains applies the Contains predicate on the "descriptions" field.
func DescriptionsContains(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescriptions), v))
	})
}

// DescriptionsHasPrefix applies the HasPrefix predicate on the "descriptions" field.
func DescriptionsHasPrefix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescriptions), v))
	})
}

// DescriptionsHasSuffix applies the HasSuffix predicate on the "descriptions" field.
func DescriptionsHasSuffix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescriptions), v))
	})
}

// DescriptionsEqualFold applies the EqualFold predicate on the "descriptions" field.
func DescriptionsEqualFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescriptions), v))
	})
}

// DescriptionsContainsFold applies the ContainsFold predicate on the "descriptions" field.
func DescriptionsContainsFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescriptions), v))
	})
}

// ThumbnailEQ applies the EQ predicate on the "thumbnail" field.
func ThumbnailEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThumbnail), v))
	})
}

// ThumbnailNEQ applies the NEQ predicate on the "thumbnail" field.
func ThumbnailNEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldThumbnail), v))
	})
}

// ThumbnailIn applies the In predicate on the "thumbnail" field.
func ThumbnailIn(vs ...string) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldThumbnail), v...))
	})
}

// ThumbnailNotIn applies the NotIn predicate on the "thumbnail" field.
func ThumbnailNotIn(vs ...string) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldThumbnail), v...))
	})
}

// ThumbnailGT applies the GT predicate on the "thumbnail" field.
func ThumbnailGT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldThumbnail), v))
	})
}

// ThumbnailGTE applies the GTE predicate on the "thumbnail" field.
func ThumbnailGTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldThumbnail), v))
	})
}

// ThumbnailLT applies the LT predicate on the "thumbnail" field.
func ThumbnailLT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldThumbnail), v))
	})
}

// ThumbnailLTE applies the LTE predicate on the "thumbnail" field.
func ThumbnailLTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldThumbnail), v))
	})
}

// ThumbnailContains applies the Contains predicate on the "thumbnail" field.
func ThumbnailContains(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldThumbnail), v))
	})
}

// ThumbnailHasPrefix applies the HasPrefix predicate on the "thumbnail" field.
func ThumbnailHasPrefix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldThumbnail), v))
	})
}

// ThumbnailHasSuffix applies the HasSuffix predicate on the "thumbnail" field.
func ThumbnailHasSuffix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldThumbnail), v))
	})
}

// ThumbnailEqualFold applies the EqualFold predicate on the "thumbnail" field.
func ThumbnailEqualFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldThumbnail), v))
	})
}

// ThumbnailContainsFold applies the ContainsFold predicate on the "thumbnail" field.
func ThumbnailContainsFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldThumbnail), v))
	})
}

// CreateDateEQ applies the EQ predicate on the "create_date" field.
func CreateDateEQ(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateDate), v))
	})
}

// CreateDateNEQ applies the NEQ predicate on the "create_date" field.
func CreateDateNEQ(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateDate), v))
	})
}

// CreateDateIn applies the In predicate on the "create_date" field.
func CreateDateIn(vs ...time.Time) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateDate), v...))
	})
}

// CreateDateNotIn applies the NotIn predicate on the "create_date" field.
func CreateDateNotIn(vs ...time.Time) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateDate), v...))
	})
}

// CreateDateGT applies the GT predicate on the "create_date" field.
func CreateDateGT(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateDate), v))
	})
}

// CreateDateGTE applies the GTE predicate on the "create_date" field.
func CreateDateGTE(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateDate), v))
	})
}

// CreateDateLT applies the LT predicate on the "create_date" field.
func CreateDateLT(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateDate), v))
	})
}

// CreateDateLTE applies the LTE predicate on the "create_date" field.
func CreateDateLTE(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateDate), v))
	})
}

// StockEQ applies the EQ predicate on the "stock" field.
func StockEQ(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStock), v))
	})
}

// StockNEQ applies the NEQ predicate on the "stock" field.
func StockNEQ(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStock), v))
	})
}

// StockIn applies the In predicate on the "stock" field.
func StockIn(vs ...int) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStock), v...))
	})
}

// StockNotIn applies the NotIn predicate on the "stock" field.
func StockNotIn(vs ...int) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStock), v...))
	})
}

// StockGT applies the GT predicate on the "stock" field.
func StockGT(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStock), v))
	})
}

// StockGTE applies the GTE predicate on the "stock" field.
func StockGTE(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStock), v))
	})
}

// StockLT applies the LT predicate on the "stock" field.
func StockLT(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStock), v))
	})
}

// StockLTE applies the LTE predicate on the "stock" field.
func StockLTE(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStock), v))
	})
}

// HasCategory applies the HasEdge predicate on the "category" edge.
func HasCategory() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CategoryTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, CategoryTable, CategoryPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCategoryWith applies the HasEdge predicate on the "category" edge with a given conditions (other predicates).
func HasCategoryWith(preds ...predicate.Category) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CategoryInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, CategoryTable, CategoryPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Product) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		p(s.Not())
	})
}
