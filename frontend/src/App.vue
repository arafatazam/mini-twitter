<template>
  <div class="container">
    <div class="columns">
      <aside class="menu column is-one-quarter">
        <ul class="menu-list">
          <li><router-link class="is-active" to="/">Home</router-link></li>
          <li><router-link to="/me">Profile</router-link></li>
          <li><router-link to="/login">Login / Signup</router-link></li>
        </ul>
      </aside>
      <router-view class="container column" />
    </div>
  </div>
</template>
