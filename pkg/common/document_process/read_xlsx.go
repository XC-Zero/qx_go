package document_process

type ReadMethod int

const (
	VERTICAL ReadMethod = iota + 1
	HORIZONTAL
)

func ReadXlsx(path, sheet string, method ReadMethod) (dataList []interface{}, err error) {

}
