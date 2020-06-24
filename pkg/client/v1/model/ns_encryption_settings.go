/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsEncryptionSettings - Group encryption settings.
// Export NsEncryptionSettingsFields for advance operations like search filter etc.
var NsEncryptionSettingsFields *NsEncryptionSettings

func init(){
		
	NsEncryptionSettingsFields= &NsEncryptionSettings{
		
	}
}

type NsEncryptionSettings struct {
   
    // Is the master key set (output only).
    
 	MasterKeySet *bool `json:"master_key_set,omitempty"`
   
    // Mode of encryption.
    
   	Mode *NsEncryptionMode `json:"mode,omitempty"`
   
    // Encryption scope.
    
   	Scope *NsEncryptionScope `json:"scope,omitempty"`
   
    // Type of encryption cipher used.
    
   	Cipher *NsEncryptionCipher `json:"cipher,omitempty"`
   
    // Is encryption active (output only).
    
 	EncryptionActive *bool `json:"encryption_active,omitempty"`
   
    // Is the master key on local or external key manager (output only).
    
   	EncryptionKeyManager *NsEncryptionKeyManager `json:"encryption_key_manager,omitempty"`
}
