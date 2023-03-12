<script lang="ts">
	import { newProjectData, newProjectSectionDataValid, newProjectShowIssues } from "../../util/stores";

	import FormField from "../form/FormField.svelte";
	import FormTextField from "../form/FormTextField.svelte";

    $: {
        $newProjectShowIssues = false;
        $newProjectSectionDataValid = ($newProjectData.jiraURL.length === 0 && $newProjectData.jiraProjectId.length === 0) || 
            ($newProjectData.jiraURL.length > 0 && $newProjectData.jiraProjectId.length > 0 && $newProjectData.jiraEmail.length > 0 && $newProjectData.jiraApiToken.length > 0);
    }

    let jiraEmailError: string = "";
    let jiraApiTokenError: string = "";
    let jiraUrlError: string = "";
    let jiraProjectIdError: string = "";
    $: {
        if($newProjectShowIssues) {
            if($newProjectData.jiraURL.length > 0 && $newProjectData.jiraProjectId.length === 0) {
                jiraProjectIdError = "Please enter your project Jira code.";
            } else {
                jiraProjectIdError = "";
            }

            if($newProjectData.jiraProjectId.length > 0 && $newProjectData.jiraURL.length === 0) {
                jiraUrlError = "Please enter your project Jira URL.";
            } else {
                jiraUrlError = "";
            }

            if($newProjectData.jiraEmail.length === 0) {
                jiraEmailError = "In order to integrate Jira to your project you need to provide Pitstop your Jira email.";
            } else {
                jiraEmailError = "";
            }

            if($newProjectData.jiraApiToken.length === 0) {
                jiraApiTokenError = "In order to integrate Jira to your project you need to provide Pitstop your Jira API token.";
            } else {
                jiraApiTokenError = "";
            }
        }
    }
</script>

<h1 class="text-xl">Jira Connection</h1>

<p class="text-sm">
	This step is optional. Click <span class="font-medium">Next</span> to skip it.
</p>

<div class="border border-slate-400 p-4 my-2 bg-slate-300 rounded-lg">
    <h2 class="text-sm font-semibold">The following settings are global - change them with caution.</h2>
    <FormTextField label="What's your Jira email?" placeholder="Email" name=jiraEmail bind:value={$newProjectData.jiraEmail} error={jiraEmailError}/>

    <FormTextField label="What's your Jira API key?" placeholder="API Key" name="jiraApiKey" bind:value={$newProjectData.jiraApiToken} error={jiraApiTokenError}/>
</div>

<p class="text-sm font-semibold pt-2">Leave the following fields blank to skip this step.</p>
    
<FormTextField label="What's the Jira project URL that you want to link to this project?" placeholder="Jira URL" name="jiraUrl" bind:value={$newProjectData.jiraURL} error={jiraUrlError}/>

<FormTextField label="What's the Jira project code that you want to link to this project?" placeholder="Project Code" name="jiraProjectId" bind:value={$newProjectData.jiraProjectId} error={jiraProjectIdError}/>
