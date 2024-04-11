import { flushPromises, mount } from '@vue/test-utils'
import { afterAll, afterEach, beforeAll, expect, test } from 'vitest'
import App from '../src/App.vue'

import { http, HttpResponse } from 'msw'
import { setupServer } from 'msw/node'

export const restHandlers = [
    http.get('/api/v1/users', ({}) => {
        return HttpResponse.json([
            { id: 'test1', name: 'John Doe' },
            { id: 'test2', name: 'Jane Doe' }
        ])
    }),
    http.put('/api/v1/todos/test1', ({}) => {
        return HttpResponse.json()
    }),
    http.get('/api/v1/todos', ({}) => {
        return HttpResponse.json([
            { id: 'test1', title: 'Todo 1', status: 'new', owner: { id: 'test1', name: 'John Doe'} },
            { id: 'test2', title: 'Todo 2', status: 'new', owner: { id: 'test1', name: 'John Doe'} }
        ])
    }),
    http.post('/api/v1/todos', ({}) => {
        return HttpResponse.json(
            { id: 'test1', title: 'New Todo', status: 'new', owner: { id: 'test1', name: 'John Doe'} }
        )
    }),
    http.post('/api/v1/users', ({}) => {
        return HttpResponse.json(
            { id: 'test3', name: 'New User' }
        )
    }),
    http.delete('/api/v1/todos/test1', ({}) => {
        return HttpResponse.json()
    }),
    http.delete('/api/v1/users/test1', ({}) => {
        return HttpResponse.json()
    }),
]
const server = setupServer(...restHandlers)
// Start server before all tests
beforeAll(() => server.listen({ onUnhandledRequest: 'error' }))
//  Close server after all tests
afterAll(() => server.close())
// Reset handlers after each test `important for test isolation`
afterEach(() => server.resetHandlers())

test('mount App', async () => {
    expect(App).toBeTruthy()
    const wrapper = mount(App)
    expect(wrapper.text()).toContain('TODOS')
    console.log(wrapper.text())
})

test('fetch users', async () => {
    const wrapper = mount(App)
    await flushPromises()
    await flushPromises()
    expect(wrapper.text()).toContain('John Doe')
    expect(wrapper.text()).toContain('Jane Doe')
    expect(wrapper.text()).toContain('Todo 1')
})

test('add user', async () => {
    const wrapper = mount(App)
    await flushPromises()
    await flushPromises()
    expect(wrapper.text()).toContain('John Doe')
    expect(wrapper.text()).toContain('Jane Doe')
    expect(wrapper.text()).toContain('Todo 1')
    await wrapper.find('#name').setValue('New User')
    await wrapper.find('#add-user-btn').trigger('click')
    await flushPromises()
    await flushPromises()
    console.log(wrapper.text())
    expect(wrapper.text()).toContain('New User')
})

test('add user by pressing enter key', async () => {
    const wrapper = mount(App)
    await flushPromises()
    await flushPromises()
    expect(wrapper.text()).toContain('John Doe')
    expect(wrapper.text()).toContain('Jane Doe')
    expect(wrapper.text()).toContain('Todo 1')
    await wrapper.find('#name').setValue('New User')
    await wrapper.find('#name').trigger('keyup.enter')
    await flushPromises()
    await flushPromises()
    console.log(wrapper.text())
    expect(wrapper.text()).toContain('New User')
})

test('delete user', async () => {
    const wrapper = mount(App)
    await flushPromises()
    await flushPromises()
    expect(wrapper.text()).toContain('John Doe')
    await wrapper.find('#delete-user-test1').trigger('click')
    await flushPromises()
    await flushPromises()
    expect(wrapper.text()).toBe('TODOSUsersJane DoeAdd UserTodosTodo 1Todo 2Add Todo')
})

test('select user', async () => {
    const wrapper = mount(App)
    await flushPromises()
    await flushPromises()
    expect(wrapper.text()).toContain('John Doe')
    await wrapper.find('#select-user-test1').trigger('click')
    await flushPromises()
    await flushPromises()
    expect(wrapper.find('#select-user-test1').classes('bg-blue-500')).toBe(true)
})

test('add todo', async () => {
    const wrapper = mount(App)
    await flushPromises()
    await flushPromises()
    expect(wrapper.text()).toContain('John Doe')
    expect(wrapper.text()).toContain('Todo 1')
    await wrapper.find('#title').setValue('New Todo')
    await wrapper.find('#add-todo-btn').trigger('click')
    await flushPromises()
    await flushPromises()
    console.log(wrapper.text())
    expect(wrapper.text()).toContain('New Todo')
})

test('delete todo', async () => {
    const wrapper = mount(App)
    await flushPromises()
    await flushPromises()
    expect(wrapper.text()).toContain('Todo 1')
    await wrapper.find('#delete-todo-test1').trigger('click')
    await flushPromises()
    await flushPromises()
    expect(wrapper.text()).toBe('TODOSUsersJohn DoeJane DoeAdd UserTodosTodo 2Add Todo')
})

test('cycle todo status', async () => {
    const wrapper = mount(App)
    await flushPromises()
    await flushPromises()
    expect(wrapper.text()).toContain('Todo 1')
    expect(wrapper.find('#select-todo-test1').classes('bg-gray-300')).toBe(true)
    await wrapper.find('#select-todo-test1').trigger('click')
    await flushPromises()
    await flushPromises()
    expect(wrapper.find('#select-todo-test1').classes('bg-yellow-500')).toBe(true)
    await wrapper.find('#select-todo-test1').trigger('click')
    await flushPromises()
    await flushPromises()
    expect(wrapper.find('#select-todo-test1').classes('bg-green-500')).toBe(true)
    await wrapper.find('#select-todo-test1').trigger('click')
    await flushPromises()
    await flushPromises()
    expect(wrapper.find('#select-todo-test1').classes('bg-gray-300')).toBe(true)
})