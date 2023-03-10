/**
 * Legal RAG classifications. One of: unknown (n/a), normal (green), warning (amber), danger (red)
 */
export type RagClassification = "unknown" | "normal" | "warning" | "danger";

/**
 * Survey factor sub-model for elements of a SurveySummary.
 */
export type SurveyFactor = {
    name: string,
    question: string,
    satisfaction: number
}

/**
 * Survey summary tile sub-model of a SoftwareProject.
 */
export type SurveySummary = {
    date: Date,
    factors: Array<SurveyFactor>,
	suggestions: Array<string>
}

/**
 * Sub-model that stores all information on the surveys undertook by this SoftwareProject.
 */
export type SoftwareSurveys = {
	client?: SurveySummary,
	team?: SurveySummary,
	health: HealthSummary
}

/**
 * Health summary tile sub-model of a SoftwareProject.
 */
export type HealthSummary = {
    status: RagClassification,
	message: string,
    issues: number
}

export type ProjectBudget = {
	budget: number,
	spend: number,
	spendOverTime: Array<number>,
	health: HealthSummary
}

/**
 * Primary model for the API request for a project.
 */
export type SoftwareProject = {
	id: number,
    name: string
    health: HealthSummary,
    surveys: SoftwareSurveys,
	budget: ProjectBudget
};

/**
 * Model that saturates the bare minimum properties needed to render the sidebar of available projects.
 */
export type SoftwareProjectSnippet = {
	id: number,
    name: string
}

/**
 * Categories available for the project management modal.
 */
export enum ProjectManageCategory {
    General,
	Budget,
	Jira,
	ClientSurvey,
	TeamSurvey,
	Settings
}

export type SurveyQuestion = {
    theme: string,
    question: string,
    enabled: boolean
}