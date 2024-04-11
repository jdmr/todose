<template>
    <div
        class="container mx-auto max-w-prose p-2 min-h-screen flex flex-col gap-4 items-center justify-center"
    >
        <h1 class="text-5xl tracking-widest">TODOS</h1>
        <div class="w-full flex flex-col gap-4">
            <h2 class="text-2xl tracking-wider">Users</h2>
            <div class="flex flex-col gap-2">
                <div
                    v-for="usr in users"
                    :key="usr.id"
                    class="border p-2 rounded text-lg tracking-wider flex justify-between items-center"
                >
                    <div class="flex items-center gap-2">
                        <button
                            :id="`select-user-${usr.id}`"
                            class="p-2 rounded-full hover:brightness-110 hover:shadow-lg focus:brightness-110 focus:shadow-lg transition-all duration-200"
                            :class="[
                                usr.id === selectedUser.id
                                    ? 'bg-blue-500 text-blue-50'
                                    : 'bg-gray-300 text-gray-800'
                            ]"
                            @click="selectUser(usr)"
                        >
                            <i class="i-mdi:check h-8 w-8"></i>
                        </button>
                        <div>{{ usr.name }}</div>
                    </div>
                    <div>
                        <button

                            :id="`delete-user-${usr.id}`"
                            class="bg-red-500 text-red-50 p-2 rounded-full hover:brightness-110 hover:shadow-lg focus:brightness-110 focus:shadow-lg transition-all duration-200"
                            @click="deleteUser(usr.id)"
                        >
                            <i class="i-mdi:trash-can h-8 w-8"></i>
                        </button>
                    </div>
                </div>
            </div>
            <div class="flex flex-col gap-1">
                <label for="name" class="text-lg tracking-wider"
                    >Add User</label
                >
                <div class="grid grid-cols-[1fr_auto] gap-2">
                    <input
                        v-model="user.name"
                        type="text"
                        id="name"
                        name="name"
                        class="border border-gray-300 rounded p-2 text-xl tracking-wider bg-gray-100 text-gray-800 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                        @keyup.enter="addUser"
                    />
                    <button
                        id="add-user-btn"
                        class="bg-blue-500 text-blue-50 p-2 rounded-full hover:brightness-110 hover:shadow-lg focus:brightness-110 focus:shadow-lg transition-all duration-200"
                        @click="addUser"
                    >
                        <i class="i-mdi:plus h-8 w-8"></i>
                    </button>
                </div>
            </div>

            <h2 class="text-2xl tracking-wider">Todos</h2>
            <div class="flex flex-col gap-2">
                <div
                    v-for="td in todos"
                    :key="td.id"
                    class="border p-2 rounded text-lg tracking-wider flex justify-between items-center"
                >
                    <div class="flex gap-2 items-center">
                        <button
                            :id="`select-todo-${td.id}`" @click="toggleStatus(td)" class="p-2 rounded-full hover:brightness-110 hover:shadow-lg focus:brightness-110 focus:shadow-lg transition-all duration-200"
                        :class="{
                            'bg-gray-300 text-gray-500': td.status === 'new',
                            'bg-yellow-500 text-yellow-50': td.status === 'started',
                            'bg-green-500 text-green-50': td.status === 'done'
                        }"
                        >
                            <i
                                :class="{
                                    'i-mdi:checkbox-blank-circle-outline':
                                        td.status === 'new',
                                    'i-mdi:clock': td.status === 'started',
                                    'i-mdi:checkbox-marked-circle': td.status === 'done'
                                }"
                                class="h-8 w-8"
                            ></i>
                        </button>
                        <div>{{ td.title }}</div>
                    </div>
                    <div>
                        <button
                            :id="`delete-todo-${td.id}`"
                            class="bg-red-500 text-red-50 p-2 rounded-full hover:brightness-110 hover:shadow-lg focus:brightness-110 focus:shadow-lg transition-all duration-200"
                            @click="deleteTodo(td.id)"
                        >
                            <i class="i-mdi:trash-can h-8 w-8"></i>
                        </button>
                    </div>
                </div>
            </div>
            <div class="flex flex-col gap-1">
                <label for="title" class="text-lg tracking-wider"
                    >Add Todo</label
                >
                <div class="grid grid-cols-[1fr_auto] gap-2">
                    <input
                        v-model="todo.title"
                        type="text"
                        id="title"
                        name="title"
                        class="border border-gray-300 rounded p-2 text-xl tracking-wider bg-gray-100 text-gray-800 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                        @keyup.enter="addTodo"
                    />
                    <button
                        id="add-todo-btn"
                        class="bg-blue-500 text-blue-50 p-2 rounded-full hover:brightness-110 hover:shadow-lg focus:brightness-110 focus:shadow-lg transition-all duration-200"
                        @click="addTodo"
                    >
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

interface Todo {
    id: string
    title: string
    status: string
    owner: User
}

const users = ref({} as User[])
const user = ref({} as User)
const todos = ref({} as Todo[])
const todo = ref({} as Todo)
const selectedUser = ref({} as User)

const addUser = async () => {
    console.log('adding user')
    user.value.id = nanoid()
    console.log('posting to /api/v1/users')
    const response = await fetch('/api/v1/users', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(user.value)
    })
    const data = await response.json()
    users.value.push(data)
    console.log('finished posting to /api/v1/users', users.value)
    user.value.name = ''
}

const deleteUser = async (id: string) => {
    await fetch(`/api/v1/users/${id}`, {
        method: 'DELETE'
    })
    users.value.splice(users.value.findIndex((user) => user.id === id), 1)
}

const selectUser = (usr: User) => {
    console.log(`Selected user: ${usr.name}`)
    selectedUser.value = usr
}

const fetchUsers = async () => {
    console.log('fetching users')
    const response = await fetch('/api/v1/users')
    console.log('assigned users')
    users.value = await response.json()
    console.log('fetched users', users.value)
}

const fetchTodos = async () => {
    const response = await fetch('/api/v1/todos')
    todos.value = await response.json()
}

const addTodo = async () => {
    todo.value.id = nanoid()
    todo.value.status = 'new'
    todo.value.owner = selectedUser.value
    const response = await fetch('/api/v1/todos', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(todo.value)
    })
    const data = await response.json()
    todos.value.push(data)
    todo.value.title = ''
}

const deleteTodo = async (id: string) => {
    await fetch(`/api/v1/todos/${id}`, {
        method: 'DELETE'
    })
    todos.value.splice(todos.value.findIndex((todo) => todo.id === id), 1)
}

const toggleStatus = async (td: Todo) => {
    console.log(`Toggling status of todo: ${td.title}`)
    switch (td.status) {
        case 'new':
            td.status = 'started'
            break
        case 'started':
            td.status = 'done'
            break
        case 'done':
            td.status = 'new'
            break
    }

    await fetch(`/api/v1/todos/${td.id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(td)
    })
}

onMounted(() => {
    fetchUsers()
    fetchTodos()
})
</script>
