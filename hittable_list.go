package main

// HITTABLE LIST //

type HittableList struct {
	hittables []Hittable
}

func (h_list *HittableList) Add(h Hittable) {
	h_list.hittables = append(h_list.hittables, h)
}

func (h_list HittableList) CheckHit(r Ray, t_min float64, t_max float64, hit_record *HitRecord) (is_hit bool) {

	temp_rec := HitRecord{
		NewPoint3(0, 0, 0),
		NewVec3(0, 0, 0),
		0,
		false,
	}

	var hit_anything bool = false
	var closest_so_far float64 = t_max

	for _, object := range h_list.hittables {
		if object.CheckHit(r, t_min, closest_so_far, &temp_rec) {
			hit_anything = true
			closest_so_far = temp_rec.t
			*hit_record = temp_rec
		}
	}
	return hit_anything
}

// HITTABLE LIST //
