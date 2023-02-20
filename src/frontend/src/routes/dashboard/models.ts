/**
 * Legal RAG classifications. One of: unknown (n/a), normal (green), warning (amber), danger (red)
 */
export type RagClassification = "unknown" | "normal" | "warning" | "danger";

/**
 * Survey factor sub-model for elements of a SurveySummary.
 */
export type SurveyFactor = {
    name: string,
    description: string,
    satisfaction: number,
    status: RagClassification
}

/**
 * Survey summary tile sub-model of a SoftwareProject.
 */
export type SurveySummary = {
    date: Date,
    factors: Array<SurveyFactor>
}

/**
 * Health summary tile sub-model of a SoftwareProject.
 */
export type HealthSummary = {
    status: RagClassification,
    issues: number
}

/**
 * Primary model for the API request for a project.
 */
export type SoftwareProject = {
	id: number,
    name: string
    health: HealthSummary,
    survey: SurveySummary
};