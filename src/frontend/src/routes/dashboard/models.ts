export type SurveyFactor = {
    name: string,
    description: string,
    satisfaction: number,
    status: string
}

export type SurveySummary = {
    date: Date,
    factors: Array<SurveyFactor>
}

export type HealthSummary = {
    status: string,
    issues: number
}

export type SoftwareProject = {
    name: string
    health: HealthSummary,
    survey: SurveySummary
};