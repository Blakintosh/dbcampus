<script lang="ts">
	export let profilePhotoUrl: string = "https://via.placeholder.com/150";
	export let name: string = "Jane Doe";
</script>

<div class="flex gap-4 items-center py-2 my-1">
	<img class="w-10 h-10 rounded-md" src={profilePhotoUrl} alt={name}/>
	<span class="text-md">{name}</span>
</div>