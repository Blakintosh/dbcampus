<script>
	import SmallButton from "../../../../common/SmallButton.svelte";
import SurveySelectionField from "../components/SurveySelectionField.svelte";

</script>
<div>
	<h3 class="font-semibold">
		Which questions do you want to ask the client?
	</h3>
	<ul class="my-2 border-t border-slate-600">
		<SurveySelectionField category="Requirements" question="I understand the latest version of the requirements." />
		<SurveySelectionField category="Vision" question="The latest version of the requirements is in line with my vision of the product." />
		<SurveySelectionField category="Management" question="I am confident in the management of the project." />
		<SurveySelectionField category="Necessity" question="The project remains high-priority to my organisation." />
		<SurveySelectionField category="Development" question="I am confident in my understanding of the development process." />
		<SurveySelectionField category="Communication" question="There is sufficient communication between me and the team." />
		<SurveySelectionField category="Timeline" question="I have no concerns about the project going off-track for time." />
		<SurveySelectionField category="Budget" question="I have no concerns about the project going off-track for budget." />
		<SurveySelectionField category="Issues" question="I have no concerns about any recent, unforeseen challenges in development." />
	</ul>

	<h3 class="font-semibold my-2">
		Ready to go?
	</h3>
	<SmallButton type="primary" label="Send Survey" />
</div>