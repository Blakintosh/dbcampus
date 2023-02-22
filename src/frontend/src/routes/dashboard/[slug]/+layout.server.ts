import type { SoftwareProject } from "../models";

export const load = (async () => {
    // Fetch data from BE
	let project: SoftwareProject = {
		id: 1,
		name: "New Exciting Project",
		health: {
			status: "warning",
			issues: 2
		},
		survey: {
			date: new Date(),
			factors: [
				{
					name: "Communication",
					description: "Description 1",
					satisfaction: 0.87,
					status: "normal"
				},
				{
					name: "No Scope Creep",
					description: "Description 3",
					satisfaction: 0.62,
					status: "warning"
				},
				{
					name: "Rewarding",
					description: "Description 3",
					satisfaction: 0.315,
					status: "danger"
				},
				{
					name: "Worthwhile",
					description: "Description 4",
					satisfaction: 0.7185,
					status: "warning"
				}
			]
		}
	};

	let projects: Array<SoftwareProject> = [
		project,
		{
			id: 2,
			name: "Foobar",
			health: {
				status: "warning",
				issues: 2
			},
			survey: {
				date: new Date(),
				factors: [
					{
						name: "Communication",
						description: "Description 1",
						satisfaction: 0.87,
						status: "normal"
					},
					{
						name: "No Scope Creep",
						description: "Description 3",
						satisfaction: 0.62,
						status: "warning"
					},
					{
						name: "Rewarding",
						description: "Description 3",
						satisfaction: 0.315,
						status: "danger"
					},
					{
						name: "Worthwhile",
						description: "Description 4",
						satisfaction: 0.7185,
						status: "warning"
					}
				]
			}
		}
	];

    return {projects: projects};
});