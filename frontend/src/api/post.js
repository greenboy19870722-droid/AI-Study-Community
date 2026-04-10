/**
 * Post API - 帖子相关接口封装
 * 对应后端 API:
 * - POST /api/post/create   - 发布帖子
 * - POST /api/post/update   - 更新帖子
 * - POST /api/post/delete   - 删除帖子
 * - GET  /api/post/detail   - 查询帖子详情
 * - GET  /api/post/list     - 分页查询帖子列表
 */
import request from './request.js'

/**
 * 发布帖子
 * @param {Object} data - 帖子数据
 * @param {string} data.title - 帖子标题
 * @param {string} data.content - 帖子内容
 * @param {number} data.authorId - 作者ID
 * @param {string} [data.tags] - 标签（逗号分隔）
 * @param {string} [data.coverImage] - 封面图片URL
 * @returns {Promise<{id: number}>} 返回新帖子ID
 */
export function createPost(data) {
  return request({
    url: '/api/post/create',
    method: 'POST',
    data,
  })
}

/**
 * 更新帖子
 * @param {Object} data - 帖子数据
 * @param {number} data.id - 帖子ID
 * @param {string} data.title - 帖子标题
 * @param {string} data.content - 帖子内容
 * @param {number} data.authorId - 作者ID
 * @param {string} [data.tags] - 标签（逗号分隔）
 * @param {string} [data.coverImage] - 封面图片URL
 * @returns {Promise<{id: number}>} 返回更新的帖子ID
 */
export function updatePost(data) {
  return request({
    url: '/api/post/update',
    method: 'POST',
    data,
  })
}

/**
 * 删除帖子
 * @param {number} id - 帖子ID
 * @returns {Promise<{success: boolean}>} 返回删除是否成功
 */
export function deletePost(id) {
  return request({
    url: '/api/post/delete',
    method: 'POST',
    data: { id },
  })
}

/**
 * 查询帖子详情
 * @param {number} id - 帖子ID
 * @returns {Promise<Object>} 返回帖子详情
 */
export function getPostDetail(id) {
  return request({
    url: '/api/post/detail',
    method: 'GET',
    params: { id },
  })
}

/**
 * 分页查询帖子列表
 * @param {Object} params - 查询参数
 * @param {number} [params.page=1] - 页码
 * @param {number} [params.pageSize=10] - 每页数量
 * @param {number} [params.authorId] - 作者ID（可选）
 * @param {string} [params.tags] - 标签筛选（可选）
 * @param {number} [params.status] - 状态筛选（可选）：0-已删除 1-正常 2-置顶 3-推荐
 * @param {string} [params.keyword] - 关键词搜索（可选）
 * @returns {Promise<Object>} 返回分页列表 { total, page, pageSize, list }
 */
export function getPostList(params) {
  return request({
    url: '/api/post/list',
    method: 'GET',
    params,
  })
}

// 导出默认对象，方便按需引入
export default {
  createPost,
  updatePost,
  deletePost,
  getPostDetail,
  getPostList,
}
