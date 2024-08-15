package go_001

// 查询数组中某个元素的下标
type Operation func(item any) bool

func FindIndex[T any](list []T, fu Operation) int {
	index := -1
	for i:= 0 ; i< len(list); i++{
		if(fu(list[i])){
			index = i
			break
		}
	}
	return index
}

// 删除数组最后一项并返回删除的元素
func Pop[T any](list *[]T)(*T){
	if(len(*list) == 0){
		return nil
	}else{
		last := (*list)[len(*list) - 1]
		*list = (*list)[:len(*list) - 1]
		return &last
	}
}