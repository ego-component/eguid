package eguid

import (
	"fmt"
	"math/rand"

	"github.com/gotomicro/ego/core/elog"
	"github.com/speps/go-hashids/v2"
)

type Component struct {
	hashIds *hashids.HashID
}

func newComponent(config *config, logger *elog.Component) *Component {
	cmp := &Component{}
	var err error
	hd := hashids.NewData()
	hd.Salt = config.Salt
	hd.MinLength = config.Length
	cmp.hashIds, err = hashids.NewWithData(hd)
	if err != nil {
		logger.Panic("guid new error", elog.FieldErr(err))
	}
	return cmp
}

func (c *Component) EncodeRandomInt64(id int64) (string, error) {
	// 给个随机数避免规律
	num := rand.Intn(10)
	return c.hashIds.EncodeInt64([]int64{id, int64(num)})
}

func (c *Component) EncodeInt64(id int64) (string, error) {
	return c.hashIds.EncodeInt64([]int64{id})
}

func (c *Component) DecodeInt64WithError(hash string) (int64, error) {
	ids, err := c.hashIds.DecodeInt64WithError(hash)
	if err != nil {
		return 0, fmt.Errorf("guid DecodeInt64WithError failed, err: %w", err)
	}
	return ids[0], nil
}
