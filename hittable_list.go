package main

// HITTABLE LIST //

type HittableList struct {
	hittables []Hittable
}

func (h_list *HittableList) Add(h Hittable) {
	h_list.hittables = append(h_list.hittables, h)
}

func (h_list HittableList) CheckHit(r Ray, i Interval, hit_record *HitRecord) (is_hit bool) {

	temp_rec := HitRecord{
		NewPoint3(0, 0, 0),
		NewVec3(0, 0, 0),
		0,
		false,
	}

	var hit_anything bool = false

	for _, object := range h_list.hittables {
		if object.CheckHit(r, i, &temp_rec) {
			hit_anything = true
			i.max = temp_rec.t
			*hit_record = temp_rec
		}
	}
	return hit_anything
}

// HITTABLE LIST //
