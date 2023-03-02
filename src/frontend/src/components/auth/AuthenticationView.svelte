<div class="grid min-h-[100vh] w-full bg-slate-100 shadow-inner">
	<div class="place-self-center justify-self-center shadow-md p-4 my-12 bg-white rounded-md">
		<div class="flex flex-col items-center m-2 lg:m-4">
			<img src="/pitstop_full.png" alt="logo" class="w-72 lg:w-96" />
            <slot name="heading"/>
		</div>
		<div class="flex flex-col items-center m-2 lg:m-4">
			<slot name="body"/>
		</div>
	</div>
</div>