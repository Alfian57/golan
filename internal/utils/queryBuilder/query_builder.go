package queryBuilder

import (
	"fmt"
	"strings"

	"github.com/Alfian57/belajar-golang/internal/dto"
)

type QueryBuilder struct {
	baseQuery   string
	whereClause []string
	args        []any
	orderBy     string
	orderType   string
}

func NewQueryBuilder(baseQuery string) *QueryBuilder {
	return &QueryBuilder{
		baseQuery: baseQuery,
		args:      make([]any, 0),
	}
}

func (qb *QueryBuilder) Where(condition string, args ...any) *QueryBuilder {
	qb.whereClause = append(qb.whereClause, condition)
	qb.args = append(qb.args, args...)
	return qb
}

func (qb *QueryBuilder) Search(column, searchTerm string) *QueryBuilder {
	if searchTerm != "" {
		qb.whereClause = append(qb.whereClause, fmt.Sprintf("%s LIKE ?", column))
		qb.args = append(qb.args, "%"+searchTerm+"%")
	}
	return qb
}

func (qb *QueryBuilder) OrderBy(column, orderType string) *QueryBuilder {
	if column == "" {
		column = "created_at" // default order by
	}
	qb.orderBy = column
	qb.orderType = strings.ToUpper(orderType)
	if qb.orderType != "ASC" && qb.orderType != "DESC" {
		qb.orderType = "ASC"
	}
	return qb
}

func (qb *QueryBuilder) Paginate(pagination dto.PaginationRequest) *QueryBuilder {
	qb.args = append(qb.args, pagination.Limit, pagination.GetOffset())
	return qb
}

func (qb *QueryBuilder) Build() (string, []interface{}) {
	query := qb.baseQuery

	if len(qb.whereClause) > 0 {
		query += " WHERE " + strings.Join(qb.whereClause, " AND ")
	}

	if qb.orderBy != "" {
		query += fmt.Sprintf(" ORDER BY %s %s", qb.orderBy, qb.orderType)
	}

	if len(qb.args) >= 2 {
		query += " LIMIT ? OFFSET ?"
	}

	return query, qb.args
}

func (qb *QueryBuilder) BuildCount(countQuery string) (string, []interface{}) {
	query := countQuery

	if len(qb.whereClause) > 0 {
		query += " WHERE " + strings.Join(qb.whereClause, " AND ")
	}

	argsWithoutPagination := qb.args
	if len(qb.args) >= 2 {
		argsWithoutPagination = qb.args[:len(qb.args)-2]
	}

	return query, argsWithoutPagination
}
