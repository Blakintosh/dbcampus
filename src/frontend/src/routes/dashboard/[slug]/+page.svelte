<script lang="ts">
	import ProjectOverview from "./_project/ProjectOverview.svelte";
	import { page } from '$app/stores';
</script>

<ProjectOverview {...$page.data.projects[0]}/>