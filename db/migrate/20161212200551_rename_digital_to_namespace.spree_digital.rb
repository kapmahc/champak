# This migration comes from spree_digital (originally 20111207121840)
class RenameDigitalToNamespace < ActiveRecord::Migration
  def change
    rename_table :digitals, :spree_digitals
    rename_table :digital_links, :spree_digital_links
  end
end
