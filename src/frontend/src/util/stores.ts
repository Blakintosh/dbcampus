
import { writable } from 'svelte/store';
import { ProjectManageCategory } from './models';

export const modalVisible = writable(true);
export const modalCategory = writable(ProjectManageCategory.General);