<script lang="ts">
    import "/src/app.css";
	import SideNav from "../../../components/dashboard/sidebar/SideNav.svelte";
	import ProjectViewLayout from "../../../components/dashboard/project/ProjectViewLayout.svelte";
	import PageHeading from "../../../components/common/PageHeading.svelte";
	import MultiStageForm from "../../../components/form/MultiStageForm.svelte";
	import FormTextField from "../../../components/form/FormTextField.svelte";
	import FormField from "../../../components/form/FormField.svelte";
	import ProjectDetails from "../../../components/newProject/ProjectDetails.svelte";
	import BudgetDetails from "../../../components/newProject/BudgetDetails.svelte";
	import ActionButton from "../../../components/common/ActionButton.svelte";
	import ConnectionsDetails from "../../../components/newProject/ConnectionsDetails.svelte";
	import type { CreateProjectData } from "../../../util/models";
	import { newProjectSectionDataValid, newProjectShowIssues, newProjectData } from "../../../util/stores";
	import DoneDetails from "../../../components/newProject/DoneDetails.svelte";
	import { fetchFromBackend } from "../../../util/backendService";
	import { goto } from "$app/navigation";

    const stages = [
        "Details",
        "Budget",
        "Jira",
        "Done"
    ];

	const stageComponents = [
		ProjectDetails,
		BudgetDetails,
		ConnectionsDetails,
        DoneDetails
	];

    let stage: number = 1;
    $: currentStage = stageComponents[Math.min(stageComponents.length, stage) - 1];

    const nextSection = () => {
        if(!$newProjectSectionDataValid) {
            $newProjectShowIssues = true;
            return;
        }
        $newProjectShowIssues = false;

        if(stage < stages.length) {
            stage++;
        }
    }

    const previousSection = () => {
        if(stage > 1) {
            stage--;
        }
    }

    let creatingProject: boolean = false;
    const createProject = async () => {
        creatingProject = true;
        stage = 5;

        const response = await fetch("/api/dashboard/createProject", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			credentials: "include",
			body: JSON.stringify($newProjectData),
		});

		if (response.status === 200) {
			// Go to dashboard... temp for now
			goto(`/dashboard/${$newProjectData.projectCode}`);
		} else {
            alert("Error sob emoji");
            creatingProject = false;
            // sob emoji here :( 
        }
    }
</script>

<svelte:head>
	<title>New Project - Pitstop Dashboard</title>
	<meta name="description" content="Project overview dashboard.">
</svelte:head>

<main class="flex max-h-full h-[100vh] flex-col lg:flex-row">
	<SideNav projects={undefined}/>

	<!-- Manage Modal -->

    <ProjectViewLayout gridClass="grid-cols-3 md:grid-cols-4 lg:grid-cols-6">
        <h1 class="col-span-6 text-3xl p-2">No project selected. That's to come...</h1>
    </ProjectViewLayout>

    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <div class="fixed w-full h-full bg-black/60 grid z-50 backdrop-blur-sm overflow-hidden">
    </div>
    <div class="fixed w-full h-full lg:min-w-[40rem] lg:w-[35%] lg:h-[75%] lg:max-h-[75%] translate-x-[-50%] left-[50%] top-[50%] translate-y-[-50%] z-50 bg-slate-200 border border-slate-300 text-black shadow-xl lg:rounded-lg flex flex-col items-stretch overflow-hidden">
        <div class="px-8 py-3 border-b border-slate-300 bg-slate-100 shrink-0">
            <PageHeading section="Dashboard" title="New Project"/>
        </div>
        <div class="flex-grow p-4 flex flex-col items-center gap-8 min-h-0">
            <div class="flex items-center justify-center mt-8">
                <MultiStageForm {stage} {stages}/>
            </div>
            <div class="flex flex-col justify-between items-center h-full w-full min-h-0">
                <div class="w-full overflow-auto flex flex-col items-center min-h-0">
                    <div class="w-full lg:w-[80%]">
                        <svelte:component this={currentStage}/>
                    </div>
                </div>
				
				<div class="flex w-full lg:w-[80%] mt-6 gap-8">
                    {#if stage > 1}
					    <ActionButton label="Go back" type="secondary" additionalClasses="shrink-0 grow" loading={false} on:click={previousSection} disabled={creatingProject} />
                    {/if}

                    {#if stage >= stages.length}
                        <ActionButton label={creatingProject ? "Creating..." : "Create project"} type="primary" additionalClasses="shrink-0 grow" loading={creatingProject} on:click={createProject} />
                    {:else}
                        <ActionButton label="Next" additionalClasses="shrink-0 grow" loading={false} on:click={nextSection} />
                    {/if}
				</div>
            </div>
        </div>
    </div>
</main>