-- name: CreateAdminPreset :exec
INSERT INTO Admin_Presets (preset_name, preset_value, slug) VALUES (?, ?, ?);

-- name: GetAdminPresetBySlug :one
SELECT id, preset_name, preset_value, slug FROM Admin_Presets WHERE slug = ?;

-- name: ListAdminPresets :many
SELECT id, preset_name, preset_value, slug FROM Admin_Presets
LIMIT ? OFFSET ?;

-- name: UpdateAdminPresetByID :exec
UPDATE Admin_Presets SET preset_name = ?, preset_value = ?, slug = ? WHERE id = ?;

-- name: TotalAdminPresetsCount :one
SELECT COUNT(*) FROM Admin_Presets;

-- name: DeleteAdminPresetByID :exec
DELETE FROM Admin_Presets WHERE id = ?;

