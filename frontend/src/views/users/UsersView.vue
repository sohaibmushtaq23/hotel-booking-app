<template>
    <v-container fluid>
      <v-row>
        <!-- Left column: user list -->
        <v-col cols="6">
          <UserList
            :users="store.users"
            :selected-user-id="store.selectedUser?.id"
            :loading="store.loading"
            :error="store.error"
            @select="store.selectUser"
            @add="openAddDialog"
            @delete="openDeleteDialog(store.selectedUser!)"
            />

        </v-col>
  
        <!-- Right column: user details -->
        <v-col cols="6">
          <UserDetails
            ref="userDetailsRef"
            :user="store.selectedUser"
            :user-roles="userRoles"
            @show-message="showNotification"
            @update-user="handleUpdateUser"
          />
        </v-col>
      </v-row>
  
      <!-- User Delete Dialog -->
      <v-dialog v-model="deleteDialog.show" max-width="400">
        <v-card title="Confirm Delete">
          <v-card-text>
            Are you sure you want to delete {{ deleteDialog.userName }}?
          </v-card-text>

          <v-card-actions>
            <v-spacer/>
            <v-btn color="grey" @click="deleteDialog.show=false">Cancel</v-btn>
            <v-btn color="success" @click="confirmDelete" :loading="store.loading">Delete</v-btn>
          </v-card-actions>
          
        </v-card>
      </v-dialog>
  
      <!-- Add User Dialog -->
      <v-dialog v-model="addDialog" max-width="600">
        <v-card title="Add New User">
          <v-card-text>
            <v-form ref="addFormRef" v-model="addValid" >
              <v-text-field v-model="newUser.userName" label="User Name" :rules="[rules.required]"></v-text-field>
              <v-text-field v-model="newUser.password" label="Password" :rules="[rules.required]" type="password"></v-text-field>
              <v-select :items="userRoles" v-model="newUser.userRole" label="User Role" :rules="[rules.required]"></v-select>
              
            </v-form>
          </v-card-text>
          <v-card-actions>
            <v-spacer/>
            <v-btn color="grey" @click="addDialog=false">Cancel</v-btn>
            <v-btn color="success" @click="saveNewUser" :loading="store.loading">Save</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
  
      <!-- Global Snackbar -->
      <v-snackbar v-model="snackbar.show" :color="snackbar.color" timeout="3000">
        {{ snackbar.message }}
        <template v-slot:actions>
          <v-btn color="white" variant="text" @click="snackbar.show = false">Close</v-btn>
        </template>
      </v-snackbar>
    </v-container>
  </template>
  
  <script setup lang="ts">
    import { ref } from 'vue'
    import { useUserStore } from '@/stores/userStore'
    import type { User } from '@/types'
    import UserList from './components/UserList.vue'
    import UserDetails from './components/UserDetails.vue'

    const store=useUserStore()
    store.fetchUsers()

    const userRoles=["Admin","Receptionist"]

    //===Add dialog===
    const addDialog=ref(false)
    const addFormRef=ref()
    const addValid=ref(false)
    const userDetailsRef = ref()

    const newUser=ref<Omit<User,'id'>>(
      {
        userName:'',
        password:'',
        userRole:''
      }
    )

    //Validation Rules
    const rules = {
      required: (v: string) => !!v || 'Required'
    }

    function openAddDialog(){
      newUser.value={
        userName:'',
        password:'',
        userRole:''
      }
      addDialog.value=true
    }

    async function saveNewUser(){
      const { valid } = await addFormRef.value.validate()
      if (!valid) return

      try{
        await store.createUser(newUser.value)
        addDialog.value=false
        showNotification('User created successfully')

      }catch(err){
        showNotification('Failed to create the user','error')

      }
    }

    //===Delete dialog===
    const deleteDialog=ref({
      show: false,
      userId:0,
      userName:''
    })

    function openDeleteDialog(user: User){

      deleteDialog.value={
        show: true,
        userId:user.id,
        userName:user.userName
      }
    }

    async function confirmDelete() {
      try{
        await store.deleteUser(deleteDialog.value.userId)
        deleteDialog.value.show=false
        showNotification('User deleted successfully')

      }catch(err){
        showNotification('Failed to delete the user','error')

      }
      
    }

    //===Update Client===
    async function handleUpdateUser(updatedUser: User) {
      try {
        await store.updateUser(updatedUser.id, updatedUser)
        showNotification('User updated successfully')
      } catch (err) {
        showNotification('Failed to update user', 'error')
      }
    }

    //===Snackbar===
    const snackbar=ref({
      show: false,
      message:'',
      color:'success'
    })

    function showNotification(message: string, color= 'success'){
      snackbar.value={show:true, message, color}
    }

  </script>