<script lang="ts">
	import ProjectSurveysView from "../_project/ProjectSurveysView.svelte";
	import { page } from '$app/stores';
</script>

<ProjectSurveysView {...$page.data.projects[0]}/>