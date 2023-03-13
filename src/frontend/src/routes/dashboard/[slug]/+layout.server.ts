import type { SoftwareProject } from "../../../util/models";
import { error, redirect } from "@sveltejs/kit";
import { goto } from "$app/navigation";

// Placeholder
const projects: Array<SoftwareProject> = [
	{
		id: 1,
		name: "New Exciting Project",
		health: {
			status: "warning",
			message: "Dk",
			issues: 2
		},
		surveys: {
			client: {
				date: new Date(),
				factors: [
					{
						name: "Communication",
						question: "Description 1",
						satisfaction: 0.87
					},
					{
						name: "No Scope Creep",
						question: "Description 3",
						satisfaction: 0.62
					},
					{
						name: "Rewarding",
						question: "Description 3",
						satisfaction: 0.315
					},
					{
						name: "Worthwhile",
						question: "Description 4",
						satisfaction: 0.7185
					}
				],
				suggestions: []
			},
			team: {
				date: new Date(),
				factors: [
					{
						name: "Communication",
						question: "Description 1",
						satisfaction: 0.87
					},
					{
						name: "No Scope Creep",
						question: "Description 3",
						satisfaction: 0.62
					},
					{
						name: "Rewarding",
						question: "Description 3",
						satisfaction: 0.315
					},
					{
						name: "Worthwhile",
						question: "Description 4",
						satisfaction: 0.7185
					}
				],
				suggestions: []
			},
			health: {
				status: "danger",
				message: "Major issues are present in the client survey.",
				issues: 4
			}
		},
		budget: {
			budget: 20000,
			spend: 12000,
			spendOverTime: [
				0, 15, 20, 27, 53, 54, 56, 60
			],
			health: {
				status: "normal",
				message: "Project is within budget.",
				issues: 0
			}
		}
	},
	{
		id: 2,
		name: "Foo",
		health: {
			message: "Placeholder",
            suggestions: [],
            percentageHealth: 0.512
		},
		surveys: {
			client: {
				date: new Date(),
				factors: [
					{
						name: "Communication",
						question: "Has the team been communicating effectively with you?",
						satisfaction: 0.87
					},
					{
						name: "Specification",
						question: "Has the team been adhering to the scope of the project?",
						satisfaction: 0.62
					},
					{
						name: "Rewarding",
						question: "Has the team been rewarding to work with?",
						satisfaction: 0.512
					},
					{
						name: "Worthwhile",
						question: "Has the team been worthwhile to work with?",
						satisfaction: 0.376
					},
					{
						name: "Budget",
						question: "Has the team been adhering to the budget?",
						satisfaction: 0.19
					}
				],
				suggestions: []
			},
			team: {
				date: new Date(),
				factors: [
					{
						name: "Communication",
						question: "Description 1",
						satisfaction: 0.87
					},
					{
						name: "No Scope Creep",
						question: "Description 3",
						satisfaction: 0.62
					},
					{
						name: "Rewarding",
						question: "Description 3",
						satisfaction: 0.315
					},
					{
						name: "Worthwhile",
						question: "Description 4",
						satisfaction: 0.7185
					}
				],
				suggestions: []
			},
			health: {
				message: "Responses with critically low satisfaction are present in the client survey.",
				suggestions: [],
                percentageHealth: 0.3
			}
		},
		budget: {
			budget: 20000,
			spend: 12345.67,
			spendOverTime: [
				0, 1058, 5056, 12120, 14150, 15650, 18650
			],
			health: {
				message: "Project is trending for completion under budget.",
				suggestions: [],
                percentageHealth: 1
			}
		}
	}
];

const getProject = (async (id: number) => {

	for(const project of projects) {
		if(project.id == id) return project;
	}

	throw error(404, "Project not found");
});

const getAvailableProjects = (async (event) => {
	const response = await event.fetch("/api/dashboard/getProjects", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
    });

    const projectsAvailable = await response.json();

    if(projectsAvailable.length === 0) {
        goto("/dashboard/new");
    }

	return await projectsAvailable;
});

// @ts-expect-error - Params any-type
export const load = (async (event) => {
	/* 
		BE connection: Fetch the project with the given slug from the database,
		return it if the user owns it and it exists, otherwise void the request.
	*/

    const currentProject = await event.fetch("/api/dashboard/getProject", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            projectCode: event.params.slug
        })
    });

    if(currentProject.status === 404) {
        throw redirect(301, "/dashboard");
    } else if(currentProject.status === 401) {
        throw redirect(301, "/auth/login");
    } else if(currentProject.status === 500) {
        throw error(500, "Unable to contact the backend. Please try again later.");
    }

    const availableProjectsResponse = await event.fetch("/api/dashboard/getProjects", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
    });

    if(availableProjectsResponse.status === 401) {
        throw redirect(301, "/auth/login");
    } else if(availableProjectsResponse.status === 500) {
        throw error(500, "Unable to contact the backend. Please try again later.");
    }

    const availableProjects = await availableProjectsResponse.json();

    if(!availableProjects || availableProjects.length === 0) {
        throw redirect(301, "/dashboard/new");
    }

    return {
		// Selected project
		project: currentProject,
		// All available projects
		availableProjects: availableProjects.data
	};
});