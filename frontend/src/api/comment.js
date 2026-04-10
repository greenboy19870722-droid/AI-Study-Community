/**
 * Comment API - 评论相关接口封装
 * 对应后端 API:
 * - POST /api/comment/create - 创建评论
 * - POST /api/comment/reply  - 回复评论
 * - POST /api/comment/delete - 删除评论
 * - GET  /api/comment/detail - 查询评论详情
 * - GET  /api/comment/tree   - 分页查询帖子下评论（层级结构）
 */
import request from './request.js'

/**
 * 创建评论
 * @param {Object} data - 评论数据
 * @param {number} data.postId - 帖子ID
 * @param {string} data.content - 评论内容
 * @param {number} data.authorId - 作者ID
 * @param {number} [data.parentId] - 父评论ID（可选，用于回复）
 * @param {number} [data.replyToUserId] - 被回复用户ID（可选）
 * @returns {Promise<{id: number}>} 返回新评论ID
 */
export function createComment(data) {
  return request({
    url: '/api/comment/create',
    method: 'POST',
    data,
  })
}

/**
 * 回复评论
 * @param {Object} data - 回复数据
 * @param {number} data.postId - 帖子ID
 * @param {number} data.parentId - 父评论ID
 * @param {string} data.content - 回复内容
 * @param {number} data.authorId - 回复者ID
 * @param {number} [data.replyToUserId] - 被回复用户ID（可选）
 * @returns {Promise<{id: number}>} 返回新回复评论ID
 */
export function replyComment(data) {
  return request({
    url: '/api/comment/reply',
    method: 'POST',
    data,
  })
}

/**
 * 删除评论
 * @param {Object} data - 删除参数
 * @param {number} data.id - 评论ID
 * @param {number} data.authorId - 作者ID（权限校验）
 * @returns {Promise<{success: boolean}>} 返回删除是否成功
 */
export function deleteComment(data) {
  return request({
    url: '/api/comment/delete',
    method: 'POST',
    data,
  })
}

/**
 * 查询评论详情
 * @param {Object} params - 查询参数
 * @param {number} params.id - 评论ID
 * @returns {Promise<Object>} 返回评论详情
 */
export function getCommentDetail(params) {
  return request({
    url: '/api/comment/detail',
    method: 'GET',
    params,
  })
}

/**
 * 分页查询帖子下评论（层级结构）
 * @param {Object} params - 查询参数
 * @param {number} params.postId - 帖子ID
 * @param {number} [params.page=1] - 页码
 * @param {number} [params.pageSize=20] - 每页数量
 * @returns {Promise<Object>} 返回评论树 { total, list, tree }
 */
export function getCommentTree(params) {
  return request({
    url: '/api/comment/tree',
    method: 'GET',
    params,
  })
}

// 导出默认对象，方便按需引入
export default {
  createComment,
  replyComment,
  deleteComment,
  getCommentDetail,
  getCommentTree,
}
