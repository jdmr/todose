<template>
    <div
        class="container mx-auto max-w-prose p-2 min-h-screen flex flex-col gap-4 items-center justify-center"
    >
        <div v-if="!authenticated.id">
            <h1 class="text-5xl tracking-widest">LOGIN</h1>
            <div class="flex flex-col gap-4">
                <div class="flex flex-col gap-2">
                    <label for="username" class="text-lg tracking-wider"
                        >Username</label
                    >
                    <input
                        v-model="authenticated.username"
                        type="text"
                        id="username"
                        name="username"
                        class="border border-gray-300 rounded p-2 text-xl tracking-wider bg-gray-100 text-gray-800 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                    />
                </div>
                <div class="flex flex-col gap-2">
                    <label for="password" class="text-lg tracking-wider"
                        >Password</label
                    >
                    <input
                        v-model="authenticated.password"
                        type="password"
                        id="password"
                        name="password"
                        class="border border-gray-300 rounded p-2 text-xl tracking-wider bg-gray-100 text-gray-800 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                    />
                </div>
                <button
                    class="bg-blue-500 text-blue-50 p-2 w-full rounded hover:brightness-110 hover:shadow-lg focus:brightness-110 focus:shadow-lg transition-all duration-200"
                    @click="login"
                >
                    <i class="i-mdi:login h-8 w-8"></i>
                </button>
            </div>
        </div>
        <div v-else>
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
                            @click="openUserDialog"
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
                                @click="toggleStatus(td)"
                                class="p-2 rounded-full hover:brightness-110 hover:shadow-lg focus:brightness-110 focus:shadow-lg transition-all duration-200"
                                :class="{
                                    'bg-gray-300 text-gray-500':
                                        td.status === 'new',
                                    'bg-yellow-500 text-yellow-50':
                                        td.status === 'started',
                                    'bg-green-500 text-green-50':
                                        td.status === 'done'
                                }"
                            >
                                <i
                                    :class="{
                                        'i-mdi:checkbox-blank-circle-outline':
                                            td.status === 'new',
                                        'i-mdi:clock': td.status === 'started',
                                        'i-mdi:checkbox-marked-circle':
                                            td.status === 'done'
                                    }"
                                    class="h-8 w-8"
                                ></i>
                            </button>
                            <div>{{ td.title }}</div>
                        </div>
                        <div>
                            <button
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
                            class="bg-blue-500 text-blue-50 p-2 rounded-full hover:brightness-110 hover:shadow-lg focus:brightness-110 focus:shadow-lg transition-all duration-200"
                            @click="addTodo"
                        >
                            <i class="i-mdi:plus h-8 w-8"></i>
                        </button>
                    </div>
                    <button
                        class="mt-4 bg-gray-500 text-gray-50 p-2 rounded-full hover:brightness-110 hover:shadow-lg focus:brightness-110 focus:shadow-lg transition-all duration-200"
                        @click="logout"
                    >
                        LOGOUT
                    </button>
                </div>
            </div>
        </div>
        <dialog ref="userDialog">
            <div class="p-2 rounded-xl flex flex-col gap-4">
                <div class="flex justify-between items-center gap-4">
                    <div class="text-lg">
                        Please provide more information to add user
                    </div>
                    <div>
                        <button
                            class="bg-red-500 text-red-50 p-2 w-full rounded-full hover:brightness-110 hover:shadow-lg focus:brightness-110 focus:shadow-lg transition-all duration-200"
                            @click="closeUserDialog"
                        >
                            <i class="i-mdi:close h-6 w-6"></i>
                        </button>
                    </div>
                </div>
                <div class="flex flex-col gap-2">
                    <label for="dialog-username" class="text-lg tracking-wider"
                        >Username</label
                    >
                    <input
                        v-model="user.username"
                        type="text"
                        id="dialog-username"
                        name="username"
                        class="border border-gray-300 rounded p-2 text-xl tracking-wider bg-gray-100 text-gray-800 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                    />
                </div>
                <div class="flex flex-col gap-2">
                    <label for="dialog-password" class="text-lg tracking-wider"
                        >Password</label
                    >
                    <input
                        v-model="user.password"
                        type="password"
                        id="dialog-password"
                        name="password"
                        class="border border-gray-300 rounded p-2 text-xl tracking-wider bg-gray-100 text-gray-800 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                    />
                </div>
                <fieldset class="flex flex-col gap-2 border rounded-lg p-2">
                    <legend>Roles</legend>
                    <label class="flex items-center gap-2">
                        <input
                            type="checkbox"
                            value="admin"
                            v-model="user.scope"
                        />
                        <span>Admin</span>
                    </label>
                    <label class="flex items-center gap-2">
                        <input
                            type="checkbox"
                            value="user"
                            v-model="user.scope"
                        />
                        <span>User</span>
                    </label>
                </fieldset>
                <button
                    class="bg-blue-500 text-blue-50 px-2 py-4 w-full rounded-lg hover:brightness-110 hover:shadow-lg focus:brightness-110 focus:shadow-lg transition-all duration-200 flex items-center justify-center gap-1"
                    @click="addUser"
                >
                    <i class="i-mdi:plus h-6 w-6"></i>
                    <span class="text-xl">Add User</span>
                </button>
            </div>
        </dialog>
    </div>
</template>

<script setup lang="ts">
import { jwtDecode } from 'jwt-decode';
import { nanoid } from 'nanoid';
import { nextTick, onMounted, ref } from 'vue';

interface User {
    id: string
    name: string
    username: string
    password: string
    scope: string[]
}

interface Todo {
    id: string
    title: string
    status: string
    owner: User
}

const users = ref({} as User[])
const user = ref({
    scope: [] as string[]
} as User)
const todos = ref({} as Todo[])
const todo = ref({} as Todo)
const selectedUser = ref({} as User)
const authenticated = ref({} as User)
let token = ''
const userDialog = ref<HTMLDialogElement | null>(null)

const login = async () => {
    console.log('logging in')
    try {
        const response = await fetch('/api/v1/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(authenticated.value)
        })
        const data = await response.json()
        token = data.token
        sessionStorage.setItem('token', token)
        const td = jwtDecode(token) as User
        authenticated.value = {
            id: td.id,
            name: td.name,
            username: td.username,
            password: '',
            scope: td.scope
        }
        fetchUsers()
        fetchTodos()
        console.log('logged in', authenticated.value)
        user.value = {
            name: '',
            username: '',
            password: '',
            scope: [] as string[]
        } as User
    } catch (error) {
        console.error('error logging in', error)
        alert('Invalid credentials')
    }
}

const openUserDialog = () => {
    console.log('opening user dialog')
    if (!userDialog.value) {
        return
    }
    if (user.value.name === '') {
        alert('Please enter a name')
        return
    }
    userDialog.value.showModal()
    nextTick(() => {
        userDialog.value?.querySelector('input')?.focus()
    })
}

const closeUserDialog = () => {
    console.log('closing user dialog')
    if (!userDialog.value) {
        return
    }
    userDialog.value.close()
}

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
    user.value = {
        scope: [] as string[]
    } as User
    userDialog.value?.close()
}

const deleteUser = async (id: string) => {
    await fetch(`/api/v1/users/${id}`, {
        method: 'DELETE'
    })
    fetchUsers()
}

const selectUser = (usr: User) => {
    console.log(`Selected user: ${usr.name}`)
    selectedUser.value = usr
}

const fetchUsers = async () => {
    console.log('fetching users')
    const response = await fetch('/api/v1/users', {
        headers: {
            Authorization: `Bearer ${token}`
        }
    })
    console.log('assigned users')
    users.value = await response.json()
    console.log('fetched users', users.value)
}

const fetchTodos = async () => {
    const response = await fetch('/api/v1/todos', {
        headers: {
            Authorization: `Bearer ${token}`
        }
    })
    todos.value = await response.json()
}

const addTodo = async () => {
    todo.value.id = nanoid()
    todo.value.status = 'new'
    todo.value.owner = selectedUser.value
    await fetch('/api/v1/todos', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(todo.value)
    })
    todo.value.title = ''
    fetchTodos()
}

const deleteTodo = async (id: string) => {
    await fetch(`/api/v1/todos/${id}`, {
        method: 'DELETE'
    })
    fetchTodos()
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

const logout = () => {
    sessionStorage.removeItem('token')
    authenticated.value = {
        id: '',
        name: '',
        username: '',
        password: '',
        scope: []
    }
    users.value = []
    todos.value = []
}

onMounted(() => {
    token = sessionStorage.getItem('token') || ''
    if (!token) {
        return
    }
    const td = jwtDecode(token) as User
    authenticated.value = {
        id: td.id,
        name: td.name,
        username: td.username,
        password: '',
        scope: td.scope
    }
    fetchUsers()
    fetchTodos()
})
</script>
