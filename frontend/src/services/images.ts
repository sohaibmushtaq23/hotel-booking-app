import api from './api'

export const uploadImage = (formData: FormData) => {
    return api.post('/upload', formData, {
        headers: { 'Content-Type': 'multipart/form-data' }
    }).then(res => res.data)
}