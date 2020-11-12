package generate

import (
	"fmt"

	"github.com/mvandergrift/energy-gql/graph/model"
)

type Collection struct {
	Query  string
	Values []interface{}
}

func buildPredicate(p *model.Predicate) string {
	ps := ""

	switch *p.Operator {
	case model.OperatorEqual:
		ps = fmt.Sprintf("%s IN (?)", *p.Name)
	case model.OperatorNotEqual:
		ps = fmt.Sprintf("%s NOT IN (?)", *p.Name)
	case model.OperatorGreator:
		ps = fmt.Sprintf("%s > ?", *p.Name)
	case model.OperatorLess:
		ps = fmt.Sprintf("%s < ?", *p.Name)
	case model.OperatorGreatorEqual:
		ps = fmt.Sprintf("%s >= ?", *p.Name)
	case model.OperatorLessEqual:
		ps = fmt.Sprintf("%s <= ?", *p.Name)
	case model.OperatorContains:
		ps = fmt.Sprintf("%s LIKE %%?%%", *p.Name)
	case model.OperatorStartsWith:
		ps = fmt.Sprintf("%s LIKE %%?", *p.Name)
	case model.OperatorIsNull:
		ps = fmt.Sprintf("%s IS NULL", *p.Name)
	case model.OperatorIsNotNull:
		ps = fmt.Sprintf("%s IS NOT NULL", *p.Name)
	}
	return ps
}

func BuildPredicateGroup(group *model.PredicateGroup, c Collection) Collection {
	for i, p := range group.Predicates {
		if p != nil {
			if p.Name != nil {
				if i > 0 {
					c.Query = fmt.Sprintf("%s %s", c.Query, group.Logic.String())
				}

				c.Query = fmt.Sprintf("%s %s", c.Query, buildPredicate((p)))
				if p.Values != nil {
					c.Values = append(c.Values, p.Values)
				}

			}

			if p.InnerPredicate != nil {
				if i > 0 {
					c.Query = fmt.Sprintf("%s %s ", c.Query, group.Logic.String())
				}

				c.Query = c.Query + "("
				c = BuildPredicateGroup(p.InnerPredicate, c)
				c.Query = c.Query + ")"
			}
		}
	}

	return c
}
