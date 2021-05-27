package job

import (
	"context"
	"fmt"
	"github.com/gotomicro/ego/task/ejob"
)

func InstallComponent() *ejob.Component {
	return ejob.DefaultContainer().Build(
		ejob.WithName("install"),
		ejob.WithStartFunc(runInstall),
	)
}

func runInstall(ctx context.Context) error {
	//models := []interface{}{
	//	&mysql.Topic{},
	//	&mysql.TopicCate{},
	//	&mysql.Abilities{},
	//}
	//gormdb := invoker.Db.Debug()
	//gormdb.SingularTable(true)
	//err := gormdb.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models...).Error
	//if err != nil {
	//	return err
	//}
	fmt.Println("create table ok")
	return nil
}
