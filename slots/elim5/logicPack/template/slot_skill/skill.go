package slot_skill

type Skill[T any] struct {
	Id        int
	SkillInfo T
}

func (s Skill[T]) GetId() int {
	return s.Id
}

func (s Skill[T]) GetValue() T {
	return s.SkillInfo
}
