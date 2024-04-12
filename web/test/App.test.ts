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
    http.get('/api/v1/todos', ({}) => {
        return HttpResponse.json([
            { id: 'test1', title: 'Todo 1', status: 'new', owner: { id: 'test1', name: 'John Doe'} },
            { id: 'test2', title: 'Todo 2', status: 'new', owner: { id: 'test1', name: 'John Doe'} }
        ])
    }),
    http.post('/api/v1/users', ({}) => {
        return HttpResponse.json(
            { id: 'test3', name: 'New User' }
        )
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
    expect(wrapper.text()).toContain('Jane Doe')
    expect(wrapper.text()).toContain('Todo 1')
    await wrapper.find('#delete-user-btn').trigger('click')
    await flushPromises()
    await flushPromises()
    console.log(wrapper.text())
    expect(wrapper.text()).not.toContain('John Doe')
})

test('select user', async () => {
    const wrapper = mount(App)
    await flushPromises()
    await flushPromises()
    expect(wrapper.text()).toContain('John Doe')
    expect(wrapper.text()).toContain('Jane Doe')
    expect(wrapper.text()).toContain('Todo 1')
    await wrapper.find('#user-select').setValue('Jane Doe')
    await flushPromises()
    await flushPromises()
    console.log(wrapper.text())
    expect(wrapper.text()).toContain('Jane Doe')
    expect(wrapper.text()).toContain('Todo 2')
})

test('fetch todos', async () => {
    const wrapper = mount(App)
    await flushPromises()
    await flushPromises()
    expect(wrapper.text()).toContain('John Doe')
    expect(wrapper.text()).toContain('Jane Doe')
    expect(wrapper.text()).toContain('Todo 1')
    expect(wrapper.text()).toContain('Todo 2')
})

test('add todo', async () => {
    const wrapper = mount(App)
    await flushPromises()
    await flushPromises()
    expect(wrapper.text()).toContain('John Doe')
    expect(wrapper.text()).toContain('Jane Doe')
    expect(wrapper.text()).toContain('Todo 1')
    await wrapper.find('#title').setValue('New Todo')
    await wrapper.find('#add-todo-btn').trigger('click')
    await flushPromises()
    await flushPromises()
    console.log(wrapper.text())
    expect(wrapper.text()).toContain('New Todo')
})

test('add todo by pressing enter key', async () => {
    const wrapper = mount(App)
    await flushPromises()
    await flushPromises()
    expect(wrapper.text()).toContain('John Doe')
    expect(wrapper.text()).toContain('Jane Doe')
    expect(wrapper.text()).toContain('Todo 1')
    await wrapper.find('#title').setValue('New Todo')
    await wrapper.find('#title').trigger('keydown.enter')
    await flushPromises()
    await flushPromises()
    console.log(wrapper.text())
    expect(wrapper.text()).toContain('New Todo')
})

test('delete todo', async () => {
    const wrapper = mount(App)
    await flushPromises()
    await flushPromises()
    expect(wrapper.text()).toContain('John Doe')
    expect(wrapper.text()).toContain('Jane Doe')
    expect(wrapper.text()).toContain('Todo 1')
    await wrapper.find('#delete-todo-btn').trigger('click')
    await flushPromises()
    await flushPromises()
    console.log(wrapper.text())
    expect(wrapper.text()).not.toContain('Todo 1')
})

test('update todo status', async () => {
    const wrapper = mount(App)
    await flushPromises()
    await flushPromises()
    expect(wrapper.text()).toContain('John Doe')
    expect(wrapper.text()).toContain('Jane Doe')
    expect(wrapper.text()).toContain('Todo 1')
    await wrapper.find('#update-todo-btn').trigger('click')
    await flushPromises()
    await flushPromises()
    console.log(wrapper.text())
    expect(wrapper.text()).toContain('done')
})

test('update todo status by pressing enter key', async () => {
    const wrapper = mount(App)
    await flushPromises()
    await flushPromises()
    expect(wrapper.text()).toContain('John Doe')
    expect(wrapper.text()).toContain('Jane Doe')
    expect(wrapper.text()).toContain('Todo 1')
    await wrapper.find('#update-todo-btn').trigger('keydown.enter')
    await flushPromises()
    await flushPromises()
    console.log(wrapper.text())
    expect(wrapper.text()).toContain('done')
})

