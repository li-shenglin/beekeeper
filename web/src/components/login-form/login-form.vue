<template>
  <div class="login">
    <Login @on-submit="handleSubmit">
      <UserName name="username" />
      <Password name="password" />
      <div class="auto-login">
        <Checkbox v-model="autoLogin" size="large">自动登录</Checkbox>
        <a>忘记密码</a>
      </div>
      <Submit />
    </Login>
  </div>
</template>
<script>
  export default {
    name: "LoginForm",
    data () {
      return {
        autoLogin: true
      }
    },
    methods: {
      handleSubmit (valid, { username, password }) {
        if (valid) {
          this.$Modal.info({
            title: '输入的内容如下：',
            content: 'username: ' + username + ' | password: ' + password
          });
        }
      }
    }
  }
</script>
<style>
  .login{
    width: 400px;
    margin: 0 auto !important;
  }
  .auto-login{
    margin-bottom: 24px;
    text-align: left;
  }
  .auto-login a{
    float: right;
  }
</style>
