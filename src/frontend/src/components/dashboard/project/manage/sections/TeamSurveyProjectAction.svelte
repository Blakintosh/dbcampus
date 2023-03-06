<script lang="ts">
	import { fetchFromBackend } from "../../../../../util/backendService";
	import type { SurveyQuestion } from "../../../../../util/models";
	import ActionButton from "../../../../common/ActionButton.svelte";
	import SmallButton from "../../../../common/SmallButton.svelte";
    import SurveySelectionField from "../components/SurveySelectionField.svelte";

    let questions: Array<SurveyQuestion> = [];

    async function getTeamSurveyQuestions() {
		const res = await fetchFromBackend("/getTeamSurveyQuestions");
		
        questions = await res.json();
        return true;
	}

    $: enabledQuestions = questions.filter(q => q.enabled).length;
	
	let promise = getTeamSurveyQuestions();
</script>
<div>
	<h3 class="font-semibold">
		Which questions do you want to ask the development team?
	</h3>
    <ul class="my-2 border-t border-b border-slate-600 max-h-[28rem] overflow-scroll">
        {#await promise}
            <li class="flex items-center text-lg p-2 gap-8 bg-slate-700 font-medium">
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="w-6 h-6 animate-spin fill-amber-300"><!--! Font Awesome Pro 6.3.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. --><path d="M304 48c0-26.5-21.5-48-48-48s-48 21.5-48 48s21.5 48 48 48s48-21.5 48-48zm0 416c0-26.5-21.5-48-48-48s-48 21.5-48 48s21.5 48 48 48s48-21.5 48-48zM48 304c26.5 0 48-21.5 48-48s-21.5-48-48-48s-48 21.5-48 48s21.5 48 48 48zm464-48c0-26.5-21.5-48-48-48s-48 21.5-48 48s21.5 48 48 48s48-21.5 48-48zM142.9 437c18.7-18.7 18.7-49.1 0-67.9s-49.1-18.7-67.9 0s-18.7 49.1 0 67.9s49.1 18.7 67.9 0zm0-294.2c18.7-18.7 18.7-49.1 0-67.9S93.7 56.2 75 75s-18.7 49.1 0 67.9s49.1 18.7 67.9 0zM369.1 437c18.7 18.7 49.1 18.7 67.9 0s18.7-49.1 0-67.9s-49.1-18.7-67.9 0s-18.7 49.1 0 67.9z"/></svg>
                    <span>Loading available survey questions...</span>
            </li>
        {:then}
            {#each questions as question}
                <SurveySelectionField theme={question.theme} question={question.question} />
            {/each}
        {:catch error}
            <li class="flex items-center text-lg p-2 gap-8 bg-slate-700 font-normal">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="h-6 w-6 fill-red-500"><!--! Font Awesome Pro 6.3.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. --><path d="M256 32c14.2 0 27.3 7.5 34.5 19.8l216 368c7.3 12.4 7.3 27.7 .2 40.1S486.3 480 472 480H40c-14.3 0-27.6-7.7-34.7-20.1s-7-27.8 .2-40.1l216-368C228.7 39.5 241.8 32 256 32zm0 128c-13.3 0-24 10.7-24 24V296c0 13.3 10.7 24 24 24s24-10.7 24-24V184c0-13.3-10.7-24-24-24zm32 224a32 32 0 1 0 -64 0 32 32 0 1 0 64 0z"/></svg>
                <span class="text-amber-300">Something went wrong. Please try again later.</span>
            </li>
        {/await}
    </ul>

    {#await promise then}
        <h3 class="font-semibold my-2">
            You are going to ask {enabledQuestions} questions. Ready to go?
        </h3>
        <ActionButton type="primary" label="Send Survey" />
    {/await}
</div>