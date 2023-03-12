<script lang="ts">
	import { newProjectData, newProjectSectionDataValid, newProjectShowIssues } from "../../util/stores";
	import FormCheckValues from "../form/FormCheckValues.svelte";

	import FormField from "../form/FormField.svelte";
	import FormTextField from "../form/FormTextField.svelte";
</script>

<h1 class="text-xl">All set?</h1>

<p class="text-sm">
    Double check all these details are correct then click <span class="font-medium">Create</span>.
</p>

<FormCheckValues values={$newProjectData} />