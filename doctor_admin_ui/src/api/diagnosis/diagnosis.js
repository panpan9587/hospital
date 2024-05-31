import service from '@/utils/request'

// @Tags Disnosis
// @Summary 连接在线问诊获取数据
// @Security ApiKeyDisnosis
// @accept application/json
// @Produce application/json
// @Param data body model.User true "连接在线问诊"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"连接成功"}"
// @Router /user/createUser [post]
export const WebDisnosis = (data) => {
  return service({
    url: '/diagnosis/WebDisnosis',
    method: 'get',
    data
  })
}
