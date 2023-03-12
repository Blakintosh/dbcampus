<script lang="ts">
	import { newProjectData, newProjectSectionDataValid, newProjectShowIssues } from "../../util/stores";
	import FormField from "../form/FormField.svelte";
	import FormTextField from "../form/FormTextField.svelte";

    let deadlineDate: string = $newProjectData.deadline.toISOString().split("T")[0];
    $: $newProjectData.deadline = new Date(deadlineDate);

    let projectNameError: string = "";
    let projectCodeError: string = "";
    let deadlineError: string = "";
    $: {
        $newProjectShowIssues = false;
        $newProjectSectionDataValid = $newProjectData.projectName.length > 0 && 
            $newProjectData.deadline > new Date() &&
            deadlineDate !== "" &&
            $newProjectData.projectCode.length === 3;
    }

    $: {
        if($newProjectShowIssues) {
            if($newProjectData.projectName.length === 0) {
                projectNameError = "Please enter a project name.";
            } else {
                projectNameError = "";
            }

            if($newProjectData.projectCode.length !== 3) {
                projectCodeError = "Project code must be 3 characters long.";
            } else {
                projectCodeError = "";
            }

            if($newProjectData.deadline < new Date() || deadlineDate === "") {
                deadlineError = "Please enter a date that is in the future.";
            } else {
                deadlineError = "";
            }
        }
    };
</script>

<h1 class="text-xl">Basic Details</h1>

<p class="text-sm">
	If you don't know what to put in here, don't worry. You can always change it later.
</p>
<p class="text-sm">
	<span class="text-red-600 font-semibold">*</span> = Mandatory
</p>

<FormTextField label="What is the project called?" placeholder="Project Name" required name="projectName" bind:value={$newProjectData.projectName} error={projectNameError}/>

<FormTextField label="What should be the project code?" placeholder="Project Code (e.g. abc)" required name="projectCode" bind:value={$newProjectData.projectCode} error={projectCodeError}/>

<FormField label="When is the project set to be finished?" isFor="finishDate" required error={deadlineError}>
	<input type="date" name="finishDate" class="w-full border {deadlineError ? "border-red-400 bg-red-100/80" : "border-slate-300"} rounded-lg p-2 mt-2" bind:value={deadlineDate}/>
</FormField>