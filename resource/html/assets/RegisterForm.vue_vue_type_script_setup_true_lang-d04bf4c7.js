var x=(y,r,i)=>new Promise((t,c)=>{var p=s=>{try{u(i.next(s))}catch(d){c(d)}},m=s=>{try{u(i.throw(s))}catch(d){c(d)}},u=s=>s.done?t(s.value):Promise.resolve(s.value).then(p,m);u((i=i.apply(y,r)).next())});import{u as z,a as I,b as R,L as F,_ as P}from"./LoginFormTitle.vue_vue_type_script_setup_true_lang-ea17ed79.js";import{S as B}from"./index-cf464bf9.js";import{C as U}from"./index-7f09b489.js";import{k as E}from"./index.js";import{$ as w,w as f,a0 as L,B as k}from"./antd-12f11a56.js";import{d as N,f as _,r as V,c as T,u as e,_ as $,$ as D,l as a,aa as n,E as g,a1 as v,ac as G}from"./vue-5c68ae35.js";const M={key:0},Q=N({__name:"RegisterForm",setup(y){const r=w.Item,i=f.Password,{t}=E(),{handleBackLogin:c,getLoginState:p}=z(),m=_(),u=_(!1),s=V({account:"",password:"",confirmPassword:"",mobile:"",sms:"",policy:!1}),{getFormRules:d}=I(s),{validForm:b}=R(m),C=T(()=>e(p)===F.REGISTER);function h(){return x(this,null,function*(){yield b()})}return(S,l)=>C.value?($(),D("div",M,[a(P,{class:"enter-x"}),a(e(w),{class:"p-4 enter-x",model:s,rules:e(d),ref_key:"formRef",ref:m},{default:n(()=>[a(e(r),{name:"account",class:"enter-x"},{default:n(()=>[a(e(f),{class:"fix-auto-fill",size:"large",value:s.account,"onUpdate:value":l[0]||(l[0]=o=>s.account=o),placeholder:e(t)("sys.login.userName")},null,8,["value","placeholder"])]),_:1}),a(e(r),{name:"mobile",class:"enter-x"},{default:n(()=>[a(e(f),{size:"large",value:s.mobile,"onUpdate:value":l[1]||(l[1]=o=>s.mobile=o),placeholder:e(t)("sys.login.mobile"),class:"fix-auto-fill"},null,8,["value","placeholder"])]),_:1}),a(e(r),{name:"sms",class:"enter-x"},{default:n(()=>[a(e(U),{size:"large",class:"fix-auto-fill",value:s.sms,"onUpdate:value":l[2]||(l[2]=o=>s.sms=o),placeholder:e(t)("sys.login.smsCode")},null,8,["value","placeholder"])]),_:1}),a(e(r),{name:"password",class:"enter-x"},{default:n(()=>[a(e(B),{size:"large",value:s.password,"onUpdate:value":l[3]||(l[3]=o=>s.password=o),placeholder:e(t)("sys.login.password")},null,8,["value","placeholder"])]),_:1}),a(e(r),{name:"confirmPassword",class:"enter-x"},{default:n(()=>[a(e(i),{size:"large",visibilityToggle:"",value:s.confirmPassword,"onUpdate:value":l[4]||(l[4]=o=>s.confirmPassword=o),placeholder:e(t)("sys.login.confirmPassword")},null,8,["value","placeholder"])]),_:1}),a(e(r),{class:"enter-x",name:"policy"},{default:n(()=>[a(e(L),{checked:s.policy,"onUpdate:checked":l[5]||(l[5]=o=>s.policy=o),size:"small"},{default:n(()=>[g(v(e(t)("sys.login.policy")),1)]),_:1},8,["checked"])]),_:1}),a(e(k),{type:"primary",class:"enter-x",size:"large",block:"",onClick:h,loading:u.value},{default:n(()=>[g(v(e(t)("sys.login.registerButton")),1)]),_:1},8,["loading"]),a(e(k),{size:"large",block:"",class:"mt-4 enter-x",onClick:e(c)},{default:n(()=>[g(v(e(t)("sys.login.backSignIn")),1)]),_:1},8,["onClick"])]),_:1},8,["model","rules"])])):G("",!0)}});export{Q as _};
