import B from"./CurrentPermissionMode-0d8fd3f5.js";import{X as P,l as T,b2 as F,_ as v}from"./index.js";import{A as g}from"./index-f1874d00.js";import{P as h}from"./index-9fe0b115.js";import{d as U,c as D,a8 as a,q as k,_ as r,a9 as n,aa as e,l as o,a0 as c,E as s,a1 as d,ac as f,n as C}from"./vue-5c68ae35.js";import{Z as $,V,a1 as N}from"./antd-12f11a56.js";import"./useContentViewHeight-c2d18fda.js";import"./useWindowSizeFn-997fa1d0.js";import"./onMountedOrActivated-97a1bb6a.js";const b=U({components:{Alert:$,PageWrapper:h,Space:V,CurrentPermissionMode:B,Divider:N,Authority:g},setup(){const{changeRole:u,hasPermission:l}=P(),E=T();return{userStore:E,RoleEnum:F,isSuper:D(()=>E.getRoleList.includes(F.SUPER)),isTest:D(()=>E.getRoleList.includes(F.TEST)),changeRole:u,hasPermission:l}}});const w={class:"mt-4"};function L(u,l,E,M,W,q){const R=a("CurrentPermissionMode"),_=a("Alert"),t=a("a-button"),A=a("Space"),i=a("Divider"),m=a("Authority"),S=a("PageWrapper"),p=k("auth");return r(),n(S,{title:"前端权限按钮示例",contentBackground:"",contentClass:"p-4",content:"由于刷新的时候会请求用户信息接口，会根据接口重置角色信息，所以刷新后界面会恢复原样，如果不需要，可以注释 src/layout/default/index内的获取用户信息接口"},{default:e(()=>[o(R),c("p",null,[s(" 当前角色: "),c("a",null,d(u.userStore.getRoleList),1)]),o(_,{class:"mt-4",type:"info",message:"点击后请查看按钮变化","show-icon":""}),c("div",w,[s(" 权限切换(请先切换权限模式为前端角色权限模式): "),o(A,null,{default:e(()=>[o(t,{onClick:l[0]||(l[0]=y=>u.changeRole(u.RoleEnum.SUPER)),type:u.isSuper?"primary":"default"},{default:e(()=>[s(d(u.RoleEnum.SUPER),1)]),_:1},8,["type"]),o(t,{onClick:l[1]||(l[1]=y=>u.changeRole(u.RoleEnum.TEST)),type:u.isTest?"primary":"default"},{default:e(()=>[s(d(u.RoleEnum.TEST),1)]),_:1},8,["type"])]),_:1})]),o(i,null,{default:e(()=>[s("组件方式判断权限(有需要可以自行全局注册)")]),_:1}),o(m,{value:u.RoleEnum.SUPER},{default:e(()=>[o(t,{type:"primary",class:"mx-4"},{default:e(()=>[s(" 拥有super角色权限可见 ")]),_:1})]),_:1},8,["value"]),o(m,{value:u.RoleEnum.TEST},{default:e(()=>[o(t,{color:"success",class:"mx-4"},{default:e(()=>[s(" 拥有test角色权限可见 ")]),_:1})]),_:1},8,["value"]),o(m,{value:[u.RoleEnum.TEST,u.RoleEnum.SUPER]},{default:e(()=>[o(t,{color:"error",class:"mx-4"},{default:e(()=>[s(" 拥有[test,super]角色权限可见 ")]),_:1})]),_:1},8,["value"]),o(i,null,{default:e(()=>[s("函数方式方式判断权限(适用于函数内部过滤)")]),_:1}),u.hasPermission(u.RoleEnum.SUPER)?(r(),n(t,{key:0,type:"primary",class:"mx-4"},{default:e(()=>[s(" 拥有super角色权限可见 ")]),_:1})):f("",!0),u.hasPermission(u.RoleEnum.TEST)?(r(),n(t,{key:1,color:"success",class:"mx-4"},{default:e(()=>[s(" 拥有test角色权限可见 ")]),_:1})):f("",!0),u.hasPermission([u.RoleEnum.TEST,u.RoleEnum.SUPER])?(r(),n(t,{key:2,color:"error",class:"mx-4"},{default:e(()=>[s(" 拥有[test,super]角色权限可见 ")]),_:1})):f("",!0),o(i,null,{default:e(()=>[s("指令方式方式判断权限(该方式不能动态修改权限.)")]),_:1}),C((r(),n(t,{type:"primary",class:"mx-4"},{default:e(()=>[s(" 拥有super角色权限可见 ")]),_:1})),[[p,u.RoleEnum.SUPER]]),C((r(),n(t,{color:"success",class:"mx-4"},{default:e(()=>[s(" 拥有test角色权限可见 ")]),_:1})),[[p,u.RoleEnum.TEST]]),C((r(),n(t,{color:"error",class:"mx-4"},{default:e(()=>[s(" 拥有[test,super]角色权限可见 ")]),_:1})),[[p,[u.RoleEnum.TEST,u.RoleEnum.SUPER]]])]),_:1})}const O=v(b,[["render",L],["__scopeId","data-v-e6134e80"]]);export{O as default};