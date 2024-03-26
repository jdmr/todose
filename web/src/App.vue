<template>
    <div
        class="container mx-auto max-w-prose p-2 min-h-screen flex flex-col gap-4 items-center justify-center"
    >
        <h1 class="text-5xl tracking-widest">TODOS</h1>
        <div class="w-full flex flex-col gap-4">
            <h2 class="text-2xl tracking-wider">Users</h2>
            <div class="flex flex-col gap-2">
                <div v-for="user in users" :key="user.id" class="border p-2 rounded text-lg tracking-wider flex justify-between items-center">
                    <div>{{ user.name }}</div>
                    <div>
                        <button class="bg-red-500 text-red-50 p-2 rounded-full hover:brightness-110 hover:shadow-lg focus:brightness-110 focus:shadow-lg transition-all duration-200" @click="deleteUser(user.id)">
                            <i class="i-mdi:trash-can h-8 w-8"></i>
                        </button>
                    </div>
                </div>
            </div>
            <div class="flex flex-col gap-1">
                <label for="name" class="text-lg tracking-wider">Add User</label>
                <div class="grid grid-cols-[1fr_auto] gap-2">
                    <input
                        v-model="user.name"
                        type="text"
                        id="name"
                        name="name"
                        class="border border-gray-300 rounded p-2 text-xl tracking-wider bg-gray-100 text-gray-800 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                        @keyup.enter="addUser"
                    />
                    <button class="bg-blue-500 text-blue-50 p-2 rounded-full hover:brightness-110 hover:shadow-lg focus:brightness-110 focus:shadow-lg transition-all duration-200" @click="addUser">
                        <i class="i-mdi:plus h-8 w-8"></i>
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { nanoid } from 'nanoid';
import { onMounted, ref } from 'vue';

interface User {
    id: string
    name: string
}

const users = ref({} as User[])
const user = ref({} as User)

const addUser = async () => {
    user.value.id = nanoid()
    await fetch('/api/v1/users', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(user.value)
    })
    user.value.name = ''
    fetchUsers()
}

const deleteUser = async (id: string) => {
    await fetch(`/api/v1/users/${id}`, {
        method: 'DELETE'
    })
    fetchUsers()
}

const fetchUsers = async () => {
    const response = await fetch('/api/v1/users')
    users.value = await response.json()
}

onMounted(fetchUsers)
</script>
