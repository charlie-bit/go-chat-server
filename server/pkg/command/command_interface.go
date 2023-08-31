package command

type ServerCommander interface {
	AddRunE()       // 新增cobra RunE方法调用
	Execute() error // cobra函数调用
}
