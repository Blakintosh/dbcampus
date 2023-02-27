<article class="p-4 shadow-inner flex-grow bg-slate-100 flex overflow-y-auto w-full min-w-0">
	<section class="flex-grow min-w-0">
        <slot name="mainContent"/>
    </section>
    <section class="flex-shrink-0 w-96">
        <slot name="sideContent"/>
    </section>
</article>