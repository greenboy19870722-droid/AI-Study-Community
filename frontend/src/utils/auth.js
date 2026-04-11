/**
 * 简易Auth工具 - 基于localStorage的用户认证状态管理
 * TODO: 后续可替换为Pinia store + 真实后端认证
 */


/**
 * 获取当前登录用户
 * @returns {Object|null} 用户信息 { id, username, ... } 或 null
 */
export function getCurrentUser() {
  try {
    const userStr = localStorage.getItem('username')
    if (!userStr) return null
    return userStr
  } catch {
    return null
  }
}

/**
 * 获取当前用户ID
 * @returns {number|null}
 */
export function getCurrentUserId() {
  return getCurrentUser()
}

/**
 * 检查是否已登录
 * @returns {boolean}
 */
export function isLoggedIn() {
  return getCurrentUserId() !== null
}




