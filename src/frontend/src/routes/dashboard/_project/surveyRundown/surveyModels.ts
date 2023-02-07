export class SurveyFactor {
	constructor(
		public readonly factor: string,
		public readonly description: string,
		public readonly satisfactionPercent: number,
		public readonly ragStatus?: string
	) {}
}