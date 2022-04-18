package grades

func init() {
	students = []Student{
		{
			ID:        1,
			FirstName: "abc",
			LastName:  "def",
			Grades: []Grade{
				{
					Title: "test1",
					Type:  Quiz,
					Score: 90,
				},
				{
					Title: "test5",
					Type:  Project,
					Score: 90,
				},
			},
		},
		{
			ID:        2,
			FirstName: "xyz",
			LastName:  "uvw",
			Grades: []Grade{
				{
					Title: "test2",
					Type:  Quiz,
					Score: 90,
				},
				{
					Title: "test3",
					Type:  Test,
					Score: 91,
				},
			},
		},
	}
}
