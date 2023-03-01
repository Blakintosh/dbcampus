<script lang="ts">
	export let gridClass: string = "";
</script>

<article class="p-4 shadow-inner flex-grow bg-slate-100 overflow-y-auto w-full min-w-0">
	<section class="min-w-0 grid {gridClass}">
        <slot/>
    </section>
</article>