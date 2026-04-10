/**
 * 简易Auth工具 - 基于localStorage的用户认证状态管理
 * TODO: 后续可替换为Pinia store + 真实后端认证
 */

const AUTH_KEY = 'auth_user'
const TOKEN_KEY = 'auth_token'

/**
 * 获取当前登录用户
 * @returns {Object|null} 用户信息 { id, username, ... } 或 null
 */
export function getCurrentUser() {
  try {
    const userStr = localStorage.getItem(AUTH_KEY)
    if (!userStr) return null
    return JSON.parse(userStr)
  } catch {
    return null
  }
}

/**
 * 获取当前用户ID
 * @returns {number|null}
 */
export function getCurrentUserId() {
  const user = getCurrentUser()
  return user?.id || null
}

/**
 * 检查是否已登录
 * @returns {boolean}
 */
export function isLoggedIn() {
  return getCurrentUserId() !== null
}

/**
 * 设置登录用户（模拟登录）
 * @param {Object} user - 用户信息
 */
export function setCurrentUser(user) {
  localStorage.setItem(AUTH_KEY, JSON.stringify(user))
}

/**
 * 清除登录状态（模拟登出）
 */
export function clearCurrentUser() {
  localStorage.removeItem(AUTH_KEY)
  localStorage.removeItem(TOKEN_KEY)
}

/**
 * 获取Token
 * @returns {string|null}
 */
export function getToken() {
  return localStorage.getItem(TOKEN_KEY)
}

/**
 * 检查用户是否有权限操作（作者ID匹配）
 * @param {number} authorId - 评论/帖子的作者ID
 * @returns {boolean}
 */
export function canOperate(authorId) {
  const currentUserId = getCurrentUserId()
  return currentUserId !== null && currentUserId === authorId
}
