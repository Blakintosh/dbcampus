<script lang="ts">
	export let date: Date = new Date();
</script>

<div class="mb-4">
	Last Budget Report: <span class="ml-1 font-medium">{
		date.toLocaleDateString("en-GB", {
			year: "numeric",
			month: "long",
			day: "numeric"
		})
		}</span>
</div>