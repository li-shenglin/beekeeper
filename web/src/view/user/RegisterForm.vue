<template>
    <div class="register">
        <div class="register-con">
            <Card icon="register-in" :title="$t('page.registerTitle')" :bordered="false">
                <div class="form-con">
                    <div class="user-register">
                        <Login ref="form" @on-submit="handleSubmit">
                            <UserName :placeholder="$t('placeholder.userName')" :rules="userNameRule" name="username"/>
                            <UserName prefix-ico="md-heart-outline" :placeholder="$t('placeholder.nickName')" :rules="nickNameRule" name="nickname"/>
                            <Poptip trigger="focus" placement="right" width="240">
                                <Password name="password" :rules="passwordRule" :placeholder="$t('placeholder.password')"
                                          @on-change="handleChangePassword"/>
                                <template #content>
                                    <div class="user-register-tip">
                                        <div class="user-register-tip-title" :class="passwordTip.class">
                                            {{ $t("page.passwordStrength") }}：{{ passwordTip.strong }}
                                        </div>
                                        <Progress :percent="passwordTip.percent" hide-info :stroke-width="6"
                                                  :stroke-color="passwordTip.color"/>
                                        <div class="user-register-tip-desc">
                                            {{ $t("page.passwordTooltip") }}
                                        </div>
                                    </div>
                                </template>
                            </Poptip>
                            <Password name="passwordConfirm" :rules="passwordConfirmRule" :placeholder="$t('placeholder.passwordConfirm')"/>
                            <Submit>{{$t("button.register")}}</Submit>
                        </Login>
                    </div>
                    <a class="register-tip">{{$t("page.toLogin")}}</a>
                </div>
            </Card>
        </div>
    </div>
</template>
<script>
    import constant from "@/config/constant"
    const $t = window.$t;
    export default {
        name: "RegisterForm",
        data() {
            const validatePassCheck = (rule, value, callback) => {
                if (value !== this.$refs.form.formValidate.password) {
                    callback(new Error($t("notice.confirmPasswordError")));
                } else {
                    callback();
                }
            };

            return {
                userNameRule: [
                    {
                        required: true, message: $t("notice.userNameRequired"), trigger: 'change'
                    },
                    {
                        min: 8, message: $t("notice.userNameTooShort"), trigger: 'change'
                    },
                    {
                        pattern: constant.regex.userName, message: $t("notice.userNameContent"), trigger: 'change'
                    }
                ],
                nickNameRule: [
                    {
                        required: true, message: $t("notice.nickNameRequired"), trigger: 'change'
                    }
                ],
                passwordRule: [
                    {
                        required: true, message: $t("notice.passwordRequired"), trigger: 'change'
                    },
                    {
                        min: 6, message: $t("notice.passwordTooShort"), trigger: 'change'
                    }
                ],
                passwordConfirmRule: [
                    {
                        required: true, message: $t("notice.confirmPasswordRequired"), trigger: 'change'
                    },
                    {validator: validatePassCheck, trigger: 'change'}
                ],
                passwordLen: 0
            }
        },
        computed: {
            passwordTip() {
                let strong = $t('notice.strong');
                let className = 'strong';
                let percent = this.passwordLen > 10 ? 10 : this.passwordLen;
                let color = '#19be6b';

                if (this.passwordLen < 6) {
                    strong = $t('notice.low');
                    className = 'low';
                    color = '#ed4014';
                } else if (this.passwordLen < 10) {
                    strong = $t('notice.medium');
                    // eslint-disable-next-line no-unused-vars
                    className = 'medium';
                    color = '#ff9900';
                }

                return {
                    strong,
                    class: 'user-register-tip-' + this.passwordLen < 6 ? 'low' : (this.passwordLen < 10 ? 'medium' : 'strong'),
                    percent: percent * 10,
                    color
                }
            }
        },
        methods: {
            handleChangePassword(val) {
                this.passwordLen = val.length;
            },
            handleSubmit(valid, {username, nickname, password}) {
                if (valid) {
                    this.$Modal.info({
                        title: '输入的内容如下：',
                        content: 'username: ' + username + ' | nickname: ' + nickname+ ' | password: ' + password
                    });
                }
            }
        }
    }
</script>
<style lang="less">
    .register {
        width: 100%;
        height: 100%;
        background-image: url('../../assets/images/login-bg.jpg');
        background-size: cover;
        background-position: center;
        position: relative;

        &-con {
            position: absolute;
            right: 160px;
            top: 50%;
            transform: translateY(-60%);
            /*width: 300px;*/

            &-header {
                font-size: 16px;
                font-weight: 300;
                text-align: center;
                padding: 30px 0;
            }

            .form-con {
                padding: 10px 0 0;
            }

            .login-tip {
                font-size: 10px;
                text-align: center;
                color: #c3c3c3;
            }
        }
    }

    .user-register {
        width: 400px;
        margin: 0 auto !important;
    }

    .user-register .ivu-poptip, .user-register .ivu-poptip-rel {
        display: block;
    }

    .user-register-tip {
        text-align: left;
    }

    .user-register-tip-low {
        color: #ed4014;
    }

    .user-register-tip-medium {
        color: #ff9900;
    }

    .user-register-tip-strong {
        color: #19be6b;
    }

    .user-register-tip-title {
        font-size: 14px;
    }

    .user-register-tip-desc {
        white-space: initial;
        font-size: 14px;
        margin-top: 6px;
    }
    .register-tip{
        display: block;
        margin-top: 14px;
    }
</style>
