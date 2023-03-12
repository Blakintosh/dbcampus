
import { writable } from 'svelte/store';
import { ProjectManageCategory } from './models';

export const modalVisible = writable(true);
export const modalCategory = writable(ProjectManageCategory.General);

export const newProjectData = writable({
    projectCode: "",
    projectName: "",
    budget: 0,
    monthlyExpenses: 0,
    customSpendings: 0,
    deadline: new Date(),
    managerExperience: 0,
    teamMeanExperience: 0,
    weeklyTeamMeetings: 0,
    clientMeetingsPerMonth: 0,
    jiraProjectId: "",
    jiraEmail: "",
    jiraApiToken: "",
    jiraURL: ""
});
export const newProjectSectionDataValid = writable(false);
export const newProjectShowIssues = writable(false);