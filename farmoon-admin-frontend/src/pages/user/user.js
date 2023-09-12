export const genderOptions = [
  {label: '男', value: 'M'},
  {label: '女', value: 'F'},
  {label: '未知', value: 'N'},
]

export const genderMap = {
  "M": "男",
  "F": "女",
  "N": "未知",
}

export const genderReverseMap = {
  "男": "M",
  "女": "F",
  "未知": "N",
}

export const roleColorMap = {
  "admin": "green",
  "editor": "orange-9",
  "visitor": "cyan",
}

export const roleOptions = [
  {label: 'admin', value: 'admin', color: roleColorMap["admin"]},
  {label: 'editor', value: 'editor', color: roleColorMap["editor"]},
  {label: 'visitor', value: 'visitor', color: roleColorMap["visitor"]},
]


