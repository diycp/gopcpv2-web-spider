package module

import "errors"

//代表未找到组件实例的错误类型。
var ErrNotFoundModuleInstance = errors.New("not found module instance")