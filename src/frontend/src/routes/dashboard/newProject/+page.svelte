<script lang="ts">
    import "/src/app.css";
	import SideNav from "../../../components/dashboard/sidebar/SideNav.svelte";
	import ProjectViewLayout from "../../../components/dashboard/project/ProjectViewLayout.svelte";
	import PageHeading from "../../../components/common/PageHeading.svelte";
	import MultiStageForm from "../../../components/form/MultiStageForm.svelte";
	import FormTextField from "../../../components/form/FormTextField.svelte";
	import FormField from "../../../components/form/FormField.svelte";

    const stages = [
        "Details",
        "Budget",
        "Connections",
        "Done"
    ];
</script>

<svelte:head>
	<title>Select a project - Pitstop Dashboard</title>
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
    <div class="fixed w-[40%] h-[75%] translate-x-[-50%] left-[50%] top-[50%] translate-y-[-50%] z-50 bg-slate-200 border border-slate-300 text-black shadow-xl rounded-lg flex flex-col items-stretch overflow-hidden">
        <div class="px-8 py-3 border-b border-slate-300 bg-slate-100 shrink-0">
            <PageHeading section="Dashboard" title="New Project"/>
        </div>
        <div class="flex-grow p-4 flex flex-col items-center gap-8">
            <div class="flex items-center justify-center mt-8">
                <MultiStageForm stage={1} {stages}/>
            </div>
            <div class="w-[80%]">
                <h1 class="text-xl">Basic Details</h1>

                <p class="text-sm">
                    If you don't know what to put in here, don't worry. You can always change it later.
                </p>
                <p class="text-sm">
                    <span class="text-red-600 font-semibold">*</span> = Mandatory
                </p>
                
                <FormTextField label="What is the project called?" placeholder="Project Name" required name="projectName"/>

                <FormField label="When is the project set to be finished?" isFor="finishDate" required>
                    <input type="date" name="finishDate" class="w-full border border-slate-300 rounded-lg p-2 mt-2"/>
                </FormField>
            </div>
        </div>
    </div>
</main>